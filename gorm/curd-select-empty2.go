package main

import (
	"fmt"

    "gorm.io/gorm"
    "errors"
    "gorm.io/driver/postgres"
    "time"
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
	fmt.Printf("1. %#v\n", db.Create(p))
	fmt.Printf("2. %#v\n", p.ID)
	db.Create(&p)
	fmt.Printf("3. %#v\n", db.Create(p))
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
    //result:= db.Where("price%20=?", 100).Select([]string{"code"}).Limit(10).First(stock)
    result:= db.Model(stock).Where("price%20=?", 100).Select([]string{"code"}).Limit(10).Scan(stock)
    err := result.Error
    if err!=nil{
        // First().Error
        fmt.Println("read Find().ErrRecordNotFound:", errors.Is(err, gorm.ErrRecordNotFound))
        fmt.Println("read empty stock string(record not found): ", err.Error())
        fmt.Println("read r.RowsAffected > 0: ", db.Model(stock).Where("price%20=?", 17).Limit(2).Scan(stock).RowsAffected > 0)
    }else{
        // Find/Scan 不再返回not found error, 建议用RowsAffected
        fmt.Println("read stock:", *stock)
        fmt.Println("read r.RowsAffected > 0: ", result.RowsAffected > 0)
    }

}

func main() {
    var err error
    //dsn:="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    dsn:="host=localhost user=role1 dbname=ahuigo sslmode=disable password= TimeZone=Asia/Shanghai"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    sqlDB, err := db.DB()
    // SetMaxIdleConns 设置空闲连接池中连接的最大数量
    sqlDB.SetMaxIdleConns(10)
    // SetMaxOpenConns 设置打开数据库连接的最大数量。
    sqlDB.SetMaxOpenConns(100)
    // SetConnMaxLifetime 设置了连接可复用的最大时间。
    sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("连接数据库失败")
	}
    defer sqlDB.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Stock{})
	createStock()
	selectStock()

}

