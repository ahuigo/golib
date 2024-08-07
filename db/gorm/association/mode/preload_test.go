package main

import (
	"testing"
	"tt"
)

func TestPreload(t *testing.T) {
	db := tt.Db
	type Order struct {
		ID     uint `gorm:"primarykey"`
		UserID uint
		Price  float64
	}
	type User struct {
		ID       uint `gorm:"primarykey"`
		Username string
		Orders   []Order
	}

	models := []interface{}{
		&User{}, &Order{},
	}
	db.Debug().Migrator().DropTable(models...)
	db.Debug().AutoMigrate(models...)

	// init
	user := User{
		Username: "ahuigo",
		Orders: []Order{
			{Price: 100},
			{Price: 200},
		},
	}
	db.Debug().Create(&user)

	// Preload Orders when find users
	users := []User{}
	db.Preload("Orders").Find(&users)
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4);

	// db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
	// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
	// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to

}
