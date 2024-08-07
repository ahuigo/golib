package curd

import (
	"fmt"
	"testing"
	"tt"
)

func TestDebug(t *testing.T) {
	type Stock struct {
		Code  string `gorm:"primary_key" `
		Price uint
	}

	var stock Stock
	tt.Db.First(&stock) // 查询id为1的product
	fmt.Println(stock)
	// refer: https://gorm.io/docs/logger.html set logger level
	tt.Db.Debug().First(&stock, "code = ?", "L1217")
	fmt.Println(stock)
}
