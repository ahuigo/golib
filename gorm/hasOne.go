
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

type CreditCard struct {
    gorm.Model
    Number string
    UID    string
}

type User struct {
    gorm.Model
    Name   string    `sql:"index"`
    CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
}

func main() {
  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  pf := fmt.Printf
  pf("'1'")
  if err != nil {
    println(err)
    fmt.Println(err)
    panic("连接数据库失败")
  }
  db.LogMode(true)
  //db.DropTable("users", "profiles")



  // 自动迁移模式
  db.AutoMigrate(&User{},&CreditCard{})
  user:=User{Model:gorm.Model{ID:2}, Name:"ahui2",}
  //user=User{}

  var card CreditCard
    // 创建
  db.Create(&User{Name: "ahui2"})
  db.Create(&CreditCard{UID: "ahui2", Number:"188-2"})

  // not work: db.Model(&user).Related(&card)
  // work: db.Model(&user).Related(&card, "CreditCard")
  db.First(&user).Related(&user.CreditCard, "CreditCard")

  pf("%+v\n", user)
  pf("%+v\n", card)



  defer db.Close()
}
