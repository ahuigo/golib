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
	db := tt.Db
	// init
	db.Migrator().DropTable(&Order{}, &User{})
	/*
		CREATE TABLE "orders" (
			"id" bigserial,PRIMARY KEY ("id"),
			"user_id" bigint,"price" decimal,
			CONSTRAINT "fk_users_orders" FOREIGN KEY ("user_id") REFERENCES "users"("uid")
		)
	*/
	db.Debug().AutoMigrate(&Order{}, &User{})
	db.Debug().Create(&User{Uid: 1, Username: "Alex1",
		Orders: []Order{
			{Price: 21},
			{Price: 22},
			{Price: 23},
			{Price: 24},
		},
	})
	db.Debug().Create(&User{Uid: 2, Username: "Alex2"})

	// Preload Orders when find users
	/**
	  SELECT * FROM "users" WHERE "users"."username" = 'Alex'
	  SELECT * FROM "orders" WHERE "orders"."user_id" = 1
	*/
	users := User{}
	tt.Db.Debug().Preload("Orders").Where(&User{Username: "Alex1"}).Find(&users)
	t.Logf("%#v\n", users)
}
