package curd

import (
	"testing"
	"tt"

	"github.com/lib/pq"
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
	这三种都会转成null:
	 1. 不能用 []string(nil)
	 2. 不能用 []string{}
	 2. 不能用 pq.StringArray(nil)
	这种才会转成空数组:
		1. 只能用 pq.StringArray{}
	********************************************/
	// 2. update empty
	addrs := pq.StringArray{}
	if err := tt.Db.Model(&Person{}).Debug().Where("username=?", p.Username).Update("addrs", addrs).Error; err != nil {
		t.Fatal(err)
	}

}
