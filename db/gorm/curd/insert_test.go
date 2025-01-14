package curd

import (
	"errors"
	"fmt"
	"testing"
	"time"
	"tt"

	"gorm.io/gorm"
)

func TestCreateBatch(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	ps := []Product{
		{Code: "L1217", Price: 17},
		{Code: "L1218", Price: 18},
		{Code: "L1219", Price: 18, Model: gorm.Model{ID: 1}},
	}
	err := tt.Db.Debug().Omit("ID").Create(&ps).Error
	// batch size 100
	// db.CreateInBatches(ps, 100)
	fmt.Printf("err:%v\n", err)
}
func TestCreate(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	p := Product{Code: "L1217", Price: 17, Model: gorm.Model{ID: 1}}
	err := tt.Db.Debug().Omit("ID").Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}

// Omit 只能有一个，它会覆盖掉之前的
func TestCreateOmits(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	p := Product{Code: "L1217", Price: 17, Model: gorm.Model{ID: 1}}
	err := tt.Db.Debug().Omit("ID").Omit("Price").Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}

// Omit 应该合并:ID,Price 或　id,price
func TestCreateOmitOne(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	p := Product{Code: "L1217", Price: 17, Model: gorm.Model{ID: 1}}
	err := tt.Db.Debug().Omit("ID,Price").Create(&p).Error
	// err := tt.Db.Debug().Omit("id,price").Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}
func TestCreateAutoID(t *testing.T) {
	type ProductID struct {
		// model外，要加：autoIncrement
		ID uint `gorm:"unique;primaryKey;autoIncrement" json:"id" form:"id"`
		//// 只能叫Model，否则会UpdatedAt/CreatedAt不自动更新
		CreatedAt time.Time  `json:"-"`
		UpdatedAt time.Time  `json:"-"`
		DeletedAt *time.Time `json:"-"`
	}
	// tt.Db.Migrator().DropTable(&ProductID{})
	tt.Db.AutoMigrate(&ProductID{})
	p := ProductID{}
	err := tt.Db.Debug().Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// u.ID = 123
	if u.ID == 123 {
		return errors.New("invalid ID")
	}
	return
}

func TestCreateHook(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	ps := []Product{
		{Code: "L1217", Price: 17},
		{Code: "L1218", Price: 18},
		{Code: "L1219", Price: 18, Model: gorm.Model{ID: 123}},
	}
	err := tt.Db.Debug().Create(&ps).Error
	// batch size 100
	// db.CreateInBatches(ps, 100)
	fmt.Printf("err:%v\n", err)
}

func TestCreateFromMap(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	err := tt.Db.Debug().Model(&Product{}).Create(map[string]interface{}{
		"Code": "jinzhu",
	}).Error
	fmt.Printf("err:%v\n", err)
}
