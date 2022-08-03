package main

import (
	"fmt"

	"github.com/ahuigo/glogger"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
)

var db *gorm.DB

type Stock struct {
	Code  string `gorm:"primary_key" `
	Price uint
	Count *uint `json:"count"  gorm:"default:2"`
}

func createStock() {
	p := Stock{Code: "L21", Price: 21}
	dbt := db.Begin()
	// no err here
	if err := dbt.Create(&p).Error; err != nil {
		glogger.Glogger.Info(err)
		dbt.Rollback()
		return
	}
	// pq: current transaction is aborted, commands ignored until end of transaction block
	p = Stock{Code: "L22", Price: 22}
	if err := dbt.Create(&p).Error; err != nil {
		glogger.Glogger.Info(err)
		dbt.Rollback()
		return
	}
	dbt.Commit()
}

func main() {
	var err error
	dsn := "host=localhost user=role1 dbname=ahuigo sslmode=disable password="
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db.LogMode(true)
	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("Db connection failed")
	}

	//db.AutoMigrate(&Stock{})
	createStock()
}
