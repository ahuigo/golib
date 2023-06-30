package curd

import (
	"fmt"
	"tt"
)

// 创建
func createStock() {
	p := Product{Code: "L1217", Price: 17}
	err := tt.Db.Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}
