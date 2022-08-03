package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
    ID   int    `gorm:"primary_key"`
    Name string
    Languages         []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    ID   int    `gorm:"primary_key"`
    Name string
    Users             []*User     `gorm:"many2many:user_languages;"`
}



func main() {
    db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
    pf := fmt.Printf
    pf("pf\n")
    if err != nil {
      println(err)
      fmt.Println(err)
      panic("连接数据库失败")
    }
    db.LogMode(true)
    //db.DropTable("users", "profiles")
    // 自动迁移模式
    db.AutoMigrate(&User{},&Language{})

    user:= User{ID:1}
    languages := []Language{}
    //等价
    //db.Model(&user).Related(&languages,  "Languages")
    db.Model(&user).Association("Languages").Append(Language{Name: "DE.............................................."}, Language{Name:"ahui"}).Find(&languages)

    a:=db.Model(&user).Association("Languages").Count()

    pf("user:%+v;\nlanguage:%+v\n", user, languages)
    pf("user:%+v;\nlanguage:%+v\n", user, a)

  defer db.Close()
}
