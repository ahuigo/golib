package main

import (
	"fmt"
	"strings"

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
    stock :=  &Stock{}
    cursor:= db.Where("price%20=?", 100).Select([]string{"code"}).Limit(10).Find(stock)
    err := cursor.Error
    fmt.Println("read Find().RecordNotFound():", cursor.RecordNotFound())
    fmt.Println("read empty stock string(record not found): ", strings.Contains(err.Error(),"record not found"))
    fmt.Println("read r.RowsAffected > 0: ", db.Where("price%20=?", 17).Select([]string{"code"}).Limit(10).Find(stock).RowsAffected > 0)

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

