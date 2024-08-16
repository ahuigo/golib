package curd

import (
	"fmt"
	"testing"
	"tt"
)

// empty error不支持ErrRecordNotFound，只能依赖: r.RowsAffected > 0
func selectScan() {
	id := 100
	stock := &Stock{}
	//result:= tt.Db.Where("price%20=?", 100).Select([]string{"code"}).Limit(10).First(stock)
	result := tt.Db.Model(stock).Where("price%20>=?", id).Select([]string{"code"}).Limit(10).Scan(stock)
	fmt.Printf("\nscan stock:%v\n", stock)
	fmt.Println("read r.RowsAffected > 0: ", result.RowsAffected > 0)
	err := result.Error
	if err != nil {
		panic(err)
	}

	stocks := []Stock{}
	result = tt.Db.Model(stock).Where("price%20>=?", id).Select([]string{"code"}).Limit(10).Scan(&stocks)
	fmt.Printf("scan stocks:%v, err:%v\n", stocks, result.Error)
}

func selectFind() {
	stock := &Stock{}
	err := tt.Db.Where("price%20>=?", 1).Select([]string{"code"}).Limit(10).Find(stock).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("stock:%v\n", stock)

	// select stocks
	stocks := []Stock{}
	err = tt.Db.Where("price%20>=?", 1).Select([]string{"code"}).Limit(10).Find(&stocks).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("stocks:%v\n", stocks)
}
func selectPluck() {
	stock := &Stock{}
	var code string
	err := tt.Db.Model(&stock).Debug().Where("price%20>=?", 1).Pluck("code", &code).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("code:%v\n", code)

}

func TestSelectFind(t *testing.T) {
	// 自动迁移模式
	tt.Db.AutoMigrate(&Product{})
	tt.Db.AutoMigrate(&Stock{})
	selectPluck()
	selectFind()
	selectScan()

}
