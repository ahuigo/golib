package curd

import (
	"testing"
	"tt"
)

func TestDelPk(t *testing.T) {
	// 自动迁移模式
	tt.Db.AutoMigrate(&Person{})

	// create
	tt.Db.Create(&Person{Name: "com", Username: "ahui"})
	p := &Person{}

	// read
	tt.Db.Find(p)
	p = &Person{}

	// read with where
	tt.Db.Where(Person{Name: "com"}).Find(p)

	// delete(自动判断主键p.Name)
	tt.Db.Unscoped().Delete(p)

}
