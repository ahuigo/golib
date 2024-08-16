package curd

import (
	"fmt"
	"testing"
	"tt"

	"github.com/lib/pq"
)

func TestInsertArray(t *testing.T) {
	tt.Db.AutoMigrate(&Person{})
	p := Person{
		Addrs: pq.StringArray{"a2", "b2"},
	}
	err := tt.Db.Debug().Omit("ID").Create(&p).Error
	fmt.Printf("addr:%v, err:%v\n", p.Addrs, err)
}
