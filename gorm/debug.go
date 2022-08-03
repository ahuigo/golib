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
	p := Stock{Code: "L12190", Price: 17}
	if db.NewRecord(p) {
		if err := db.Create(&p).Error; err != nil {
			panic(err)
		}
	}
}

// update
func updateStock() {
	p := Stock{Code: "L1217", Price: 19}
	db.Model(&p).Update(&p)
	//db.Model(&product).Update("Price", 22)
}

// read
func selectStock() {
	var stock Stock
	db.First(&stock) // 查询id为1的product
	fmt.Println(stock)
	db.First(&stock, "code = ?", "L1217")
	fmt.Println(stock)
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
	db.LogMode(true)
	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&Stock{})
	createStock()
	updateStock()
	selectStock()

	defer db.Close()
}

/*

 */
