package curd

import (
	"fmt"
	"testing"
	"tt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// update
func updateStock() {
	println("update L1 to code=4")
	p := Stock{Code: "L1", Price: 4}
	err := tt.Db.Model(&p).Updates(&p)
	println(err)
	//tt.Db.Model(&product).Update("Price", 22)
}

// read raw
func selectStock() {
	var stock Stock
	tt.Db.First(&stock, "code = ?", "L1217")
	fmt.Println(stock)

	//可以是指针数组
	// stocks :=  []*Stock{}
	type S struct{ Code string }
	stocks := []*S{}
	tt.Db.Raw("select * from stocks limit 2").Scan(&stocks)
	fmt.Println("read stock:", *stocks[0])

	//也可以纯指针
	sp := &S{}
	tt.Db.Raw("select * from stocks limit 2").Scan(sp)
	fmt.Println("read stockp:", *sp)

}

func TestCurd(t *testing.T) {
	// 自动迁移模式
	tt.Db.AutoMigrate(&Product{})
	tt.Db.AutoMigrate(&Stock{})
	tt.Db.AutoMigrate(&User{})
	db := tt.Db
	createStock()
	selectStock()
	updateStock()

	user := User{
		Model: gorm.Model{
			ID: 2,
		},
	}
	// update delete
	db.Debug().Model(&user).Update("Age", 18)
	db.Debug().Model(&user).Omit("Age").Updates(map[string]interface{}{"UserName": "jinzhu", "Age": 19})

	// delete
	db.Debug().Unscoped().Delete(&user)
}
