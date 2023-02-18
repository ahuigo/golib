package t

import (
	"fmt"
	"testing"

	"github.com/ahuigo/glogger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// var db *gorm.DB

func TestInsertOnColumn(t *testing.T) {
	type Stock struct {
		Code  string `gorm:"primary_key" `
		Price uint
		Count *uint `json:"count"  gorm:"default:2"`
	}
	createStock := func() {
		p := Stock{Code: "L21", Price: 21}
		// no err here(Warnning!!!!)
		if err := db.Create(&p).Error; err != nil {
			glogger.Info(err)
			return
		}
	}

	var err error
	db, err = gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
	db.LogMode(true)
	if err != nil {
		println(err)
		println(err.Error())
		fmt.Println(err)
		panic("Db connection failed")
	}

	//db.AutoMigrate(&Stock{})
	createStock()

	defer db.Close()
}
