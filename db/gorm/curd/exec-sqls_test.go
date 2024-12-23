package curd

import (
	"fmt"
	"testing"
	"tt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func TestExecSqls(t *testing.T) {
	db := tt.Db
	db.AutoMigrate(&Product{})

	err := db.Debug().Session(&gorm.Session{SkipDefaultTransaction: true}).Exec(`
    -- create database db2; -- comment
    insert into products("code") values('a2');
    insert into products("code") values('a3');
    insert into products("code") values('a4');
    `).Error
	if err != nil {
		fmt.Println("1. err:", err.Error())
	}
	var results []Product
	db.Raw("SELECT * from products").Scan(&results)
	for _, r := range results {
		fmt.Println("7. code:", r.Code)
	}
}

func TestSelectAs(t *testing.T) {
	tt.Db.Select("*, 'testMoreInfoVal' AS more_info").Where(&User{}).Find(&[]User{})

	// stocks :=  []*Stock{}
	// type S struct{ Code string }
	stocks := []*Stock{}
	tt.Db.Raw("select id as cid from stocks limit 2").Where("code", 2).Scan(&stocks)
	fmt.Println("read stock:", *stocks[0])
}
func TestWhereRaw(t *testing.T) {
	// gorm 不会支持：db.Where(&Stock{Num: 3000}).Raw("") // where 被raw覆盖
	db := tt.Db
	sql := fmt.Sprintf("select * from users where value=%s", pq.QuoteLiteral("a'b"))
	db.Raw(sql).Scan(&[]User{})
}
