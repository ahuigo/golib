package tc

import (
	"testing"
	"tt"

	"github.com/ahuigo/glogger"
)

// var db *gorm.DB

func TestAborted(t *testing.T) {
	type Stock struct {
		Code  string `gorm:"primary_key" `
		Price uint
		Count *uint `json:"count"  gorm:"default:2"`
	}
	db := tt.Db
	db.AutoMigrate(&Stock{})
	createStock := func() {
		p := Stock{Code: "L21", Price: 21}
		dbt := db.Begin()
		// no err here
		if err := dbt.Create(&p).Error; err != nil {
			glogger.Info(err)
			dbt.Rollback()
			return
		}
		// pq: current transaction is aborted, commands ignored until end of transaction block
		p = Stock{Code: "L22", Price: 22}
		if err := dbt.Create(&p).Error; err != nil {
			glogger.Info(err)
			dbt.Rollback()
			return
		}
		dbt.Commit()
	}

	//db.AutoMigrate(&Stock{})
	createStock()

}
