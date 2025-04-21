package curd

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Person struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" form:"name"`
	Username string `gorm:"uniqueIndex:idx_username" json:"username"`
	Age      int
	Valid    *bool
	Addrs    pq.StringArray `gorm:"not null;default:array[]::varchar[];type:text[]" json:"addrs" form:"addrs"`
	// Addrs3   pq.StringArray `gorm:"not null;type:text[]" json:"addrs3" form:"addrs3"`
	// insert　err: []string type is record: VALUES (('a1','b1'))
	// Addrs2   []string       `gorm:"not null;default:array[]::varchar[];type:text[]" json:"addrs2" form:"addrs2"`
}

type Product struct {
	//// 只能叫Model，否则会UpdatedAt/CreatedAt不自动更新
	gorm.Model
	Code  string
	Price uint
	Data  []byte
}

type Stock struct {
	Code      string `gorm:"primary_key" `
	Price     uint
	Num       int
	Count     *uint     `json:"count"  gorm:"default:2"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex:idx_username" json:"username"`
	Age      uint
	Score    uint
}
