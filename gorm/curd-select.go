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
func create() {
	p := Product{Code: "L1217", Price: 17}
	fmt.Printf("1. %#v\n", db.NewRecord(p))
	fmt.Printf("2. %#v\n", p.ID)
	db.Create(&p)
	fmt.Printf("3. %#v\n", db.NewRecord(p))
	fmt.Printf("4. %#v\n", p.ID)
}

// 创建
func createStock() {
	p := Stock{Code: "L1218", Price: 17}
	db.Create(&p)
	db.Create(&Stock{Code: "L1219", Price: 19})
}

// read
func selectStock() {
    stocks :=  []*Stock{}
    stock :=  Stock{}
    err := db.Where("price%21=?", 0).Select([]string{"code"}).Limit(10).Order("code desc").First(&stock).Error

    err = db.Where("price%20=?", 0).Select([]string{"code"}).Limit(10).Order("code desc").Find(&stocks).Error
    fmt.Println("read stocks:", len(stocks), err)

    // err = db.Where(nil).Limit(10).Find(&stocks).Error // invalid nil
    err = db.Where(&Stock{Code:"L1218"}).Limit(10).Find(&stocks).Error // invalid nil
    fmt.Println("read stocks:", len(stocks), err)

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

	// 自动迁移模式
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Stock{})
	createStock()
	selectStock()

	defer db.Close()
}

