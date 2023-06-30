package curd

import (
	"gorm.io/gorm"
)

type Person struct {
	Name     string `gorm:"primary_key" json:"name" form:"name"`
	Username string `gorm:"unique_index:idx_username" json:"username"`
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
