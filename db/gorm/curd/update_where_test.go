package curd

import (
	"testing"
	"tt"
)

func TestUpdateEmptyWhere(t *testing.T) {
	// 不为空的primary_key 才会被自动加入where条件. 如果where条件为空，会报错WHERE conditions required
	type Person1 struct {
		ID uint `gorm:"primary_key" json:"id"`
		// Username string `gorm:"uniqueIndex:idx_username" json:"username"`
		Username string `gorm:"primary_key;not null;" `
		Age      int
	}
	var err error
	tt.Db.Debug().Migrator().DropTable(&Person1{})
	tt.Db.Debug().AutoMigrate(&Person1{})

	// 1. insert
	p := Person1{ID: 1, Username: "Alex", Age: 4}
	if err = tt.Db.Debug().Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2. update table
	db := tt.Db.Debug()
	p.ID = 0
	// 2.1 从model中获取primary key: WHERE username='u1' AND "id" = 1 AND "username" = 'Alex'
	if err := db.Model(&p).Updates(&p).Error; err != nil {
		t.Fatal(err)
	}

}
func TestUpdateWhere(t *testing.T) {
	// 自动迁移模式
	// CREATE TABLE "person1" ("id" bigserial,"username" text NOT NULL,"age" bigint,PRIMARY KEY ("id","username"))
	type Person1 struct {
		ID uint `gorm:"primary_key" json:"id"`
		// Username string `gorm:"uniqueIndex:idx_username" json:"username"`
		Username string `gorm:"primary_key;not null;" `
		Age      int
	}
	var err error
	tt.Db.Debug().Migrator().DropTable(&Person1{})
	tt.Db.Debug().AutoMigrate(&Person1{})

	// 1. insert
	p := Person1{ID: 1, Username: "Alex", Age: 4}
	if err = tt.Db.Debug().Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2. update table
	db := tt.Db.Debug()
	// 2.1 从model中获取primary key: WHERE username='u1' AND "id" = 1 AND "username" = 'Alex'
	if err := db.Model(&p).Where("username=?", "u1").Updates(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2.2 从model清除where primary key: WHERE username='u1'
	if err := db.Model(&Person1{}).Where("username=?", "u1").Updates(&p).Error; err != nil {
		t.Fatal(err)
	}
	// 2.3 默认使用更新数据做model: WHERE "id" = 1 AND "username" = 'Alex'
	if err := db.Updates(&p).Error; err != nil {
		t.Fatal(err)
	}
	// 2.4 不会更新任何数据：WHERE conditions required
	if err := db.Updates(&Person1{Age: 30}).Error; err != nil {
		t.Fatal(err)
	}

}
