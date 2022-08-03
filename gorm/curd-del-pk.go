package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)


type Person struct {
    Name string `gorm:"primary_key" json:"name" form:"name"`
	Username string `gorm:"unique_index:idx_username" json:"username"`
}

func main() {
  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  if err != nil {
    println(err)
    println(err.Error())
    fmt.Println(err)
    panic("连接数据库失败")
  }

  // 创建
  db.LogMode(true)

  // 自动迁移模式
  db.AutoMigrate(&Person{})

  // create
  db.Create(&Person{Name:"com", Username:"ahui"})
  p := &Person{}

  // read
  db.Find(p)
  p = &Person{}

  // read with where
  db.Where(Person{Name:"com"}).Find(p)

  // delete(自动判断主键p.Name)
  db.Unscoped().Delete(p)



  defer db.Close()
}
