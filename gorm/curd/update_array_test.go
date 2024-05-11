package curd

import (
	"testing"
	"tt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func TestUpdateArray(t *testing.T) {
	// 自动迁移模式
	var err error
	tt.Db.Migrator().DropTable(&Person{})
	tt.Db.AutoMigrate(&Person{})

	// 1. insert
	p := Person{Name: "com", Username: "Alex", Age: 3, Addrs: []string{"a", "b"}}
	if err = tt.Db.Debug().Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2. find
	res := Person{}
	if err := tt.Db.Where("username=?", p.Username).Find(&res).Error; err != nil {
		panic(err)
	} else {
		t.Logf("res:%v\n", res)
		if res.Addrs[0] != "a" {
			panic("error addrs")
		}
	}

	/************* update **
	以下这三种都会转成null, 所以:
	 1. 不能用 []string(nil)
	 2. 不能用 []string{}
	 2. 不能用 pq.StringArray(nil)
	这种才会转成空数组, 只能用真正的空数组:
		1. pq.StringArray{}
		2. gorm.Expr("ARRAY[]::varchar[]")
	********************************************/
	// 2. update empty
	addrs := pq.StringArray{}
	db := tt.Db.Model(&Person{}).Debug().Where("username=?", p.Username)
	if err := db.Update("addrs", addrs).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Update("addrs", gorm.Expr("ARRAY[]::varchar[]")).Error; err != nil {
		t.Fatal(err)
	}

}
