package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Stock struct {
	Code      string `gorm:"primary_key" `
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func getDb() (db *gorm.DB, err error) {
	conn := `host=localhost user=role1 dbname=ahuigo sslmode=disable password=`
	println(conn)
	db, err = gorm.Open("postgres", conn)
	db.LogMode(true)
	if err != nil {
		err = fmt.Errorf("gorm.Open %s, err: %v", conn, err)
		println(err)
		return
	}
	return
}

func insertHandler(c *gin.Context) {
	res := "ok"
	_, err := getDb()
	if err != nil {
		res = err.Error()
	}

	c.String(http.StatusOK, res)

}
