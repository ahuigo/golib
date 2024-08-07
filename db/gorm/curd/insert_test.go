package curd

import (
	"errors"
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	p := Product{Code: "L1217", Price: 17, Model: gorm.Model{ID: 1}}
	err := tt.Db.Debug().Omit("ID").Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}

func TestCreateList(t *testing.T) {
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
