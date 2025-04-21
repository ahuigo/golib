package curd

import (
	"testing"
	"tt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// 2. update:  UPDATE "person1" SET "updated_at"='2025-04-21 21:54:49.763',"age"=5 WHERE id=1 AND "person1"."deleted_at" IS NULL AND "id" = 1
	p = Person1{ID: 1, Age: 5}
	db := tt.Db.Debug()
	if err := db.Model(&p).Where("id=?", 1).Updates(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 3. updateAll:UPDATE "person1" SET "updated_at"='2025-04-21 21:54:49.763',"age"=5 WHERE id=1 AND "person1"."deleted_at" IS NULL AND "id" = 1
	p = Person1{ID: 1, Age: 6}
	if err := tt.Db.Debug().Model(&p).
		Clauses(clause.OnConflict{
			// Columns:   []clause.Column{{Name: "xxx"}}, //指定联合主键, 默认是primary key(仅限单主键)
			UpdateAll: true,
		}).
		Where("id=?", 1).Updates(&p).Error; err != nil {
		t.Fatal(err)
	}

}
