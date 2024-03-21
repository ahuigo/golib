package main

import (
	"testing"
	"tt"

	"gorm.io/gorm"
)

func TestMigrate2(t *testing.T) {
	type CellPolygon struct {
		gorm.Model
		Code string
		// alter table products add column price serial;
		// 相当于
		// CREATE SEQUENCE mytable_item_id_seq OWNED BY mytable.item_id;
		// ALTER TABLE mytable ALTER item_id SET DEFAULT nextval('mytable_item_id_seq');
		Price uint `gorm:"AUTO_INCREMENT"`
	}
	db := tt.Db
	println("migrate begin")
	err := db.AutoMigrate(
		&CellPolygon{}, //flag1ZZ
	)
	if err != nil {
		panic(err)
	}
}
