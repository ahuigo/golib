package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type YourModel struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

func main() {
	dsn := "host=localhost user=role1 password='' dbname=ahuigo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var results []YourModel

	rawQuery := `
		WITH t(id) AS (VALUES(?),(?),(?))
		SELECT * from t;
	`

	id1 := 1
	id2 := 2
	id3 := 3

	if err := db.Raw(rawQuery, id1, id2, id3).Scan(&results).Error; err != nil {
		fmt.Println("Error executing query:", err)
	} else {
		fmt.Println("Results:", results)
	}
}
