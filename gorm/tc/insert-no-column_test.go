package tc

import (
	"testing"
	"tt"

	"github.com/ahuigo/glogger"
)

// var db *gorm.DB

func TestInsertOnColumn(t *testing.T) {
	type Stock struct {
		Code  string `gorm:"primary_key" `
		Price uint
		Count *uint `json:"count"  gorm:"default:2"`
	}
	db := tt.Db
	createStock := func() {
		p := Stock{Code: "L21", Price: 21}
		// no err here(Warnning!!!!)
		if err := db.Create(&p).Error; err != nil {
			glogger.Info(err)
			return
		}
	}

	//db.AutoMigrate(&Stock{})
	createStock()

}
