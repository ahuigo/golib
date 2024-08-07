package curd

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Person struct {
    // primary_key, index, uniqueIndex, ...
    //Name2 string `gorm:"uniqueIndex"`
    //Name2 string `gorm:"uniqueIndex:idx_name,sort:desc"`
	Name     string `gorm:"primary_key" json:"name" form:"name"`
	Username string `gorm:"uniqueIndex:idx_username" json:"username"`
	Age      int
	Valid    *bool
	Addrs    pq.StringArray `gorm:"null;default:array[]::varchar[];type:text[]" json:"addrs" form:"addrs"`
	// Addrs []string `gorm:"not null;default:array[]::varchar[];type:text[]" json:"addrs" form:"addrs"`
}

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
