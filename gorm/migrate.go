package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

type User struct {
  gorm.Model
  UserName string
  Age uint
}

func main() {
  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  if err != nil {
    println(err)
    println(err.Error())
    fmt.Println(err)
    panic("连接数据库失败")
  }


  // 自动迁移模式
  db.AutoMigrate(&Product{})

  // 创建
  db.Create(&Product{Code: "L1217", Price: 17})
  db.Create(&Product{Code: "L1218", Price: 18})

  // 读取
  var product Product
  db.First(&product, 1) // 查询id为1的product
fmt.Println(product)
  db.First(&product, "code = ?", "L1214") // 查询code为l1212的product
fmt.Println(product)

  // 更新 - 更新product的price为2000
  db.Model(&product).Update("Price", 22)

  defer db.Close()
}
