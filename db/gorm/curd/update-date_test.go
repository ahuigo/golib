package curd

import (
	"testing"
	"tt"

	"gorm.io/gorm"
)

func TestUpdateCurrentDate(t *testing.T) {
	type Person1 struct {
		gorm.Model
		ID  uint `gorm:"primary_key" json:"id"`
		Age int
	}
	var err error
	tt.Db.Debug().Migrator().DropTable(&Person1{})
	tt.Db.Debug().AutoMigrate(&Person1{})

	// 1. insert
	p := Person1{ID: 1, Age: 4}
	if err = tt.Db.Debug().Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2. update:  UPDATE "person1" SET "created_at"='2024-07-12 22:57:40.442',"updated_at"='2024-07-12 22:57:40.444',"age"=4 WHERE id=1
	db := tt.Db.Debug()
	if err := db.Model(&p).Where("id=?", 1).Updates(&p).Error; err != nil {
		t.Fatal(err)
	}

}
