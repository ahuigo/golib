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

// 创建
func create() {
    err:=db.Exec(`
    create database db2;//error
    insert into products("code") values('a2');
    insert into products("code") values('a3');
    insert into products("code") values('a4');
    `)
    fmt.Println("1. err:", err)
    var results []Product
    db.Raw("SELECT * from products").Scan(&results)
    for _,r:=range results{
        fmt.Println("7. code:", r.Code)
    }
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

	db.AutoMigrate(&Product{})
    create()

	defer db.Close()
}

