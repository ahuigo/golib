package curd

import (
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
)

func TestCreateList(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	ps := []Product{
		{Code: "L1217", Price: 17},
		{Code: "L1218", Price: 18},
		{Code: "L1219", Price: 18, Model: gorm.Model{ID: 1}},
	}
	err := tt.Db.Debug().Omit("ID").Create(&ps).Error
	// batch size 100
	// db.CreateInBatches(ps, 100)
	fmt.Printf("err:%v\n", err)
}
