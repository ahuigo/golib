package curd

import (
	"fmt"
	"strings"
	"testing"
	"tt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db1 *gorm.DB

// 创建
func insertStock() {
	p := Product{Code: "L1217", Price: 17}
	err := tt.Db.Create(&p).Error
	fmt.Printf("id:%v, err:%v\n", p.ID, err)
}

// read
func selectFindEmpty() {
	stock := &Stock{}
	cursor := db1.Where("price%20>=?", 100).Select([]string{"code"}).Limit(10).Find(stock)
	fmt.Println("stock:", stock)

	err := cursor.Error
	if err != nil {
		fmt.Println("read Find().RecordNotFound():", cursor.RecordNotFound())
		fmt.Println("read empty stock string(record not found): ", strings.Contains(err.Error(), "record not found"))
		fmt.Println("read r.RowsAffected > 0: ", cursor.RowsAffected > 0)
	}

}

func TestSelectEmpty1(t *testing.T) {
	var err error
	db1, err = gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
	db1.LogMode(true)
	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db1.AutoMigrate(&Product{})
	db1.AutoMigrate(&Stock{})
	insertStock()
	selectFindEmpty()

	defer db1.Close()
}
