package main

import (
	"testing"
	"tt"
)

// User has many Orders
func TestHasmany(t *testing.T) {
	type Order struct {
		ID     int `gorm:"primarykey"`
		UserID int
		Price  float64
	}
	type User struct {
		ID       int `gorm:"primarykey"`
		Username string
		Orders   []Order `gorm:"foreignKey:user_id;references:id;"`
	}
	/*
		CREATE TABLE "orders" ("id" bigserial,"user_id" bigint,"price" decimal,PRIMARY KEY ("id"),
		 		CONSTRAINT "fk_users_orders" FOREIGN KEY ("user_id") REFERENCES "users"("id")
		 )
	*/
	db := tt.Db
	// init
	db.Migrator().DropTable(&Order{}, &User{})
	db.Debug().AutoMigrate(&Order{}, &User{})
	db.Create(&User{ID: 1, Username: "Alex"})
	db.Create(&User{ID: 3, Username: "Alex3"})
	db.Create(&Order{ID: 1, UserID: 3, Price: 20})
	db.Create(&Order{ID: 2, UserID: 3, Price: 21})

	// Preload Orders when find users
	/**
	  SELECT * FROM "users" WHERE "users"."username" = 'Alex3'
	  SELECT * FROM "orders" WHERE "orders"."user_id" = 3
	*/
	users := User{}
	tt.Db.Debug().Preload("Orders").Where(&User{Username: "Alex3"}).Find(&users)
	t.Logf("%#v\n", users)

	// db.Migrator().DropTable(&Order{}, &User{})
}
