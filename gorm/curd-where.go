package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Stock struct {
	Code  string `gorm:"primary_key" `
	Price uint
}

type User struct {
	gorm.Model
	UserName string
	Age      uint
}

// 创建
func createStock() {
	p := Stock{Code: "L1218", Price: 17}
	db.Create(&p)
	db.Create(&Stock{Code: "L1219", Price: 19})
}

// update
func updateStock() {
	p := Stock{Code: "L1217", Price: 19}
	db.Model(&p).Where(&p).Find(&p)
	db.Model(&p).Where("code=?","1").Where("price>? or price<?", 1,100).Update(&p)
}

func main() {
    var err error
	db, err = gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
	db.LogMode(true)
	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("连接数据库失败")
	}

	updateStock()

	defer db.Close()
}

/*

 */
