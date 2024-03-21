package main

import (
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	// alter table products add column price serial;
	// 相当于
	// CREATE SEQUENCE mytable_item_id_seq OWNED BY mytable.item_id;
	// ALTER TABLE mytable ALTER item_id SET DEFAULT nextval('mytable_item_id_seq');
	Price uint `gorm:"AUTO_INCREMENT"`
}

func TestMigrate(t *testing.T) {
	db := tt.Db
	// 自动迁移模式
	// db.DropTableIfExists(&Product{})
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1217", Price: 17})
	db.Create(&Product{Code: "L1218", Price: 18})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	fmt.Println(product)
	db.First(&product, "code = ?", "L1214") // 查询code为l1212的product
	fmt.Println(product)

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 22)

}
