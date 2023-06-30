package curd

import (
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
)

func TestExecSqls(t *testing.T) {
	db := tt.Db
	db.AutoMigrate(&Product{})

	err := db.Debug().Session(&gorm.Session{SkipDefaultTransaction: true}).Exec(`
    -- create database db2; -- comment
    insert into products("code") values('a2');
    insert into products("code") values('a3');
    insert into products("code") values('a4');
    `).Error
	if err != nil {
		fmt.Println("1. err:", err.Error())
	}
	var results []Product
	db.Raw("SELECT * from products").Scan(&results)
	for _, r := range results {
		fmt.Println("7. code:", r.Code)
	}
}
