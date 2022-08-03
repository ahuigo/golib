package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "time"
)

type User struct {
    ID                uint       `gorm:"primary_key"`
    Birthday          string     `gorm:"column:birthday"`
    Age               int        `gorm:"column:age"`
    Name              string     `gorm:"column:name"`
    BillingAddress    Address    `gorm:"foreignkey:BillingAddressId;"` // One-To-One (属于 - 本表的BillingAddressID作外键
    BillingAddressId  int        `gorm:"column:billing_address_id"`
    ShippingAddress   Address    `gorm:"foreignkey:ShippingAddressId;"` // One-To-One (属于 - 本表的ShippingAddressID作外键)
    ShippingAddressId int        `gorm:"column:shipping_address_id"`
    CreditCard        CreditCard `gorm:"foreignkey:UserID;"` // One-To-One (拥有一个 - CreditCard表的UserID作外键)
    Emails            []Email    `gorm:"ForeignKey:UserID;"`            // One-To-Many (拥有多个 - Email表的UserID作外键)
    Languages         []Language `gorm:"many2many:user_languages;"`     // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
    ID     int    `gorm:"primary_key"`
    UserID int    `gorm:"column:user_id;"` // 外键 (属于), tag `index`是为该列创建索引
    Email  string `gorm:"column:email"`    // `type`设置sql类型, `unique_index` 为该列设置唯一索引
}

type Address struct {
    ID       int    `gorm:"primary_key"`
    Address1 string `gorm:"column:address1"` // 设置字段为非空并唯一
    Address2 string `gorm:"column:address2"`
    Post     string `gorm:"column:post"`
}

type Language struct {
    ID   int    `gorm:"primary_key"`
    Name string `gorm:"column:name"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
    Code string `gorm:"column:code"` // `unique_index` also works
}

type CreditCard struct {
    ID     int    `gorm:"primary_key"`
    UserID int    `gorm:"column:user_id;"`
    Number string `gorm:"column:number"`
}

func main() {
    db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
    db.LogMode(true)
    if err != nil {
        panic("failed to connect database")
    }
    db.DB().SetMaxIdleConns(30)
    db.DB().SetMaxOpenConns(60)
    defer db.Close()
    defer func() {
        if e := recover(); e != nil {
            fmt.Println(e)
        }
    }()
    t1 := time.Now()
    u := User{ID: 1}
    db.Model(&u).Find(&u)
    db.Model(&u).Related(&u.CreditCard,"CreditCard")
    db.Model(&u).Related(&u.Emails,"Emails")
    db.Model(&u).Related(&u.Languages, "Languages")
    db.Model(&u).Related(&u.BillingAddress,"BillingAddress")
    db.Model(&u).Related(&u.ShippingAddress,"ShippingAddress")
    fmt.Println(u)
    t2 := time.Now()
    fmt.Println("gorm耗时:", t2.Sub(t1))
}

func (u User) TableName() string {
    return "user"
}

func (u Email) TableName() string {
    return "email"
}

func (u Address) TableName() string {
    return "address"
}

func (u Language) TableName() string {
    return "language"
}

func (u CreditCard) TableName() string {
    return "creditcard"
}
