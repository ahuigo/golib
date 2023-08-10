package main

import (
	"testing"
	"tt"
)

// User has many Orders
func TestHasmanyCustomId(t *testing.T) {
	type Order struct {
		ID     int `gorm:"primarykey"`
		UserID int `gorm:"index"`
		Price  float64
	}
	type User struct {
		ID       int `gorm:"primarykey"`
		Uid      int `gorm:"uniqueIndex"`
		Username string
		Orders   []Order `gorm:"foreignKey:user_id;references:uid;"` // orders(user_id) has one user, user has many orders
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
	db.Create(&User{ID: 1, Uid: 1, Username: "Alex1"})
	db.Create(&User{ID: 2, Uid: 2, Username: "Alex2"})
	db.Create(&Order{ID: 1, UserID: 1, Price: 20})
	db.Create(&Order{ID: 2, UserID: 1, Price: 21})

	// Preload Orders when find users
	/**
	  SELECT * FROM "users" WHERE "users"."username" = 'Alex'
	  SELECT * FROM "orders" WHERE "orders"."user_id" = 1
	*/
	users := User{}
	tt.Db.Debug().Preload("Orders").Where(&User{Username: "Alex1"}).Find(&users)
	t.Logf("%#v\n", users)

	// db.Migrator().DropTable(&Order{}, &User{})
}
