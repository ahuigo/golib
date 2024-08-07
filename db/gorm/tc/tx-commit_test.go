package tc

import (
	"testing"
	"tt"

	"github.com/ahuigo/glogger"
)

func TestTxCommit(t *testing.T) {
	type Stock struct {
		Code  string `gorm:"primary_key" `
		Price uint
		Count *uint `json:"count"  gorm:"default:2"`
	}
	db := tt.Db
	db.AutoMigrate(&Stock{})
	createStock := func() {
		p := Stock{Code: "L21", Price: 21}
		// 1. begin
		tx := db.Begin()
		// 2. create 1
		if err := tx.Create(&p).Error; err != nil {
			glogger.Info(err)
			tx.Rollback()
			return
		}
		// 2. create 2
		p = Stock{Code: "L22", Price: 22}
		if err := tx.Create(&p).Error; err != nil {
			glogger.Info(err)
			tx.Rollback()
			return
		}

		// 3. commit or rollback
		tx.Commit()
	}

	//db.AutoMigrate(&Stock{})
	createStock()
}
