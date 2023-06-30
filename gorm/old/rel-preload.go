package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)
type User struct {
    //gorm.Model
    ID   int    `gorm:"primary_key"`
    Name string
    Languages         []Language `gorm:"many2many:user_languages;"`
    CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
}
type CreditCard struct {
    gorm.Model
    Number string
    UID    string
}

type Language struct {
    ID   int    `gorm:"primary_key"`
    Name string
    Users             []User     `gorm:"many2many:user_languages;association_jointable_foreignkey:user_id;jointable_foreignkey:language_id;foreignkey:ID"` 
                /*
                HasMany:
                    association_foreignkey:ID;//(User.ID);
                    foreignkey:ID; //(Language.ID);
                */
}


func main() {
  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  pf := fmt.Printf
  pf("'pf init'")
  if err != nil {
    println(err)
    fmt.Println(err)
    panic("连接数据库失败")
  }
  db.LogMode(true)
  //db.DropTable("users", "languages","user_languages")
  // 自动迁移模式
  db.AutoMigrate(&User{},&Language{},&CreditCard{})

    user := User{
        ID:1,
        Name:            "jinzhu",
        Languages:       []Language{
            {Name: "ZH"},
            {Name: "EN"},
        },
        CreditCard: CreditCard{UID:"jinzhu", Number:"6222"},
    }
    //db.Create(&user)
    //db.Set("gorm:association_autoupdate", false).Create(&user)
    db.Set("gorm:association_save_reference", false).Save(&user)
    //db.Save(&user)


    //create: HasMany/hasOne/Many2Many
    var users []User
    var language Language

    //Query Related
    pf("--------------\n\n")
    pf("Related Users:...\n")
    user=User{}
    db.First(&language, "id = ?", 1)
    db.Model(&language).Related(&users,  "Users")
    pf("users:%+v;\n", users)
    pf("--------------\n\n")

    //Query Preload
    pf("Preload Users:...\n")
    user=User{}
    db.Preload("Languages").Preload("CreditCard").First(&user)
    pf("user:%+v;\n\n", user, )

  defer db.Close()
}
