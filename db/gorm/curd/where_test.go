package curd

import (
	"testing"
	"time"
	"tt"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	cond := Stock{
		Code: "L1217", Price: 19,
		CreatedAt: time.Now(),
	}
	t0 := time.Now().Add(-time.Hour)
	dest := Stock{Code: "L1218", Price: 20}
	// 1. where: bind conds: &cond + &dest.PK 共同作为查询条件(忽略0、""等默认值)
	tt.Db.Debug().Model(&Stock{}).Where("created_at>?", t0).Find(&dest, &cond)
	// // 2. where: query fields
	err := tt.Db.Debug().Model(&Stock{}).Where("code=?", "1").Where("price>? or price<?", 1, 100).Updates(&cond).Error
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

func TestWhereGroupOr(t *testing.T) {
	db := tt.Db.Debug()
	// Complex SQL query using Group Conditions
	db.Where(
		"id", 1,
	).Where(
		db.Where("pizza = ?", "hawaiian").Or("size = ?", "xlarge"),
	).Find(&[]User{})
}
