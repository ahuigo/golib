package curd

import (
	"testing"
	"tt"
)

func TestWhere(t *testing.T) {
	p := Stock{Code: "L1217", Price: 19}
	p2 := Stock{Code: "L1218", Price: 20}
	tt.Db.Debug().Model(&p).Find(&p2, &p)
	tt.Db.Debug().Model(&p).Where("code=?", "1").Where("price>? or price<?", 1, 100).Updates(&p)
}
