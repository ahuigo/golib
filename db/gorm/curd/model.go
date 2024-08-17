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
	// insert　err: []string type is record: VALUES (('a1','b1'))
	// Addrs2   []string       `gorm:"not null;default:array[]::varchar[];type:text[]" json:"addrs2" form:"addrs2"`
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Stock struct {
	Code  string `gorm:"primary_key" `
	Price uint
	Num   int
	Count *uint `json:"count"  gorm:"default:2"`
}

type User struct {
	gorm.Model
	UserName string
	Age      uint
}
