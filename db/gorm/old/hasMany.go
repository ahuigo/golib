package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)


type User struct {
    gorm.Model
  MemberNumber string
    CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard struct {
    gorm.Model
    Number           string
  UserMemberNumber string
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

    var cards []CreditCard
    user:=User{Model:gorm.Model{ID:23}, MemberNumber:"ahui",}
    db.Create(&user)
    db.Create(&CreditCard{UserMemberNumber:"ahui",})
    db.Find(&user).Related(&cards, "CreditCards")//.First(&user)
    pf("user:%+v;\ncards:%+v\n", user, cards)
    db.First(&user).Related(&user.CreditCards, "CreditCards")//.First(&user)
    pf("user:%+v;\ncards:%+v\n", user, cards)

  defer db.Close()
}
