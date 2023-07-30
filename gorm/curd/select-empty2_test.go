package curd

import (
	"fmt"
	"strings"
	"testing"
	"tt"

	"errors"

	"gorm.io/gorm"
)

/**
https://stackoverflow.com/questions/68810961/gorm-use-the-find-rerurn-empty
only First, Take, Last could return ErrRecordNotFound:
	tx.Statement.RaiseErrorOnNotFound = true
*/

func selectScanEmpty2() {
	stock := &Stock{}
	result := tt.Db.Model(stock).Where("price%20=?", 100).Select([]string{"code"}).Limit(10).Scan(stock)
	fmt.Printf("scan empty: err=%v, stock=%v\n", result.Error, stock)
	err := result.Error
	if err != nil {
		panic(err)
	} else {
		// Find/Scan 不再返回not found error, 建议用RowsAffected
		fmt.Println("scan stock:", *stock)
		fmt.Println("scan r.RowsAffected > 0: ", result.RowsAffected > 0)
	}
}

func selectFindEmpty2() {
	stock := &Stock{}
	// gorm2 只有在你使用 First、Last、Take 这些预期会返回结果的方法查询记录时，才会返回 ErrRecordNotFound，我们还移除了 RecordNotFound
	cursor := tt.Db.Where("price%20>=?", 100).Select([]string{"code"}).Limit(10).Find(stock)
	err := cursor.Error
	fmt.Printf("find empty: err=%v, stock=%v\n", err, stock)
	fmt.Printf("read r.RowsAffected == 0: %v\n\n", cursor.RowsAffected == 0)
	if err != nil {
		fmt.Println("read Find().RecordNotFound():", errors.Is(err, gorm.ErrRecordNotFound))                          // otgorm
		fmt.Println("read empty stock string(record not found): ", strings.Contains(err.Error(), "record not found")) // ag -F 'record not found'
	}

}

func TestSelectEmpty2(t *testing.T) {
	var err error
	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("连接数据库失败")
	}
	defer tt.SqlDb.Close()

	// 自动迁移模式
	tt.Db.AutoMigrate(&Product{})
	tt.Db.AutoMigrate(&Stock{})
	createStock()
	selectScanEmpty2()
	selectFindEmpty2()

}
