package curd

import (
	"testing"
	"time"
	"tt"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	p := Stock{Code: "L1217", Price: 19}
	p2 := Stock{Code: "L1218", Price: 20}
	tt.Db.Debug().Model(&p).Find(&p2, &p)
	err := tt.Db.Debug().Model(&p).Where("code=?", "1").Where("price>? or price<?", 1, 100).Updates(&p).Error
	assert.NoError(t, err)
}

func TestWhereDate(t *testing.T) {
	p := struct {
		StartTime time.Time `gorm:"start_time"`
	}{}
	db := tt.Db.Debug()
	// t1 := time.Now()
	t2 := time.Now().AddDate(0, 0, 1)
	db = db.Raw("SELECT * FROM date_trunc('day', ?::timestamp) AS start_time ", t2)
	db = db.Find(&p)
	err := db.Error
	assert.NoError(t, err)
}
