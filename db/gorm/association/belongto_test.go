package main

import (
	"testing"
	"tt"
)

/*
// User has one Card
*/
func TestBelongToCard(t *testing.T) {
	type CreditCard struct {
		ID     int `gorm:"primarykey"`
		Number string
		UserID uint `gorm:"index"`
	}
	type User struct {
		ID         int `gorm:"primarykey"`
		Username   string
		CreditCard CreditCard `gorm:"foreignKey:user_id;references:id;"` //可省略, 等价于hasOne: Card(user_id) has one User(id)
	}
	/*
		CREATE TABLE "credit_cards" (
			"id" bigserial, PRIMARY KEY ("id"),
			"number" text,
			"user_id" bigint,
			CONSTRAINT "fk_users_credit_card" FOREIGN KEY ("user_id") REFERENCES "users"("id"))
		# \d credit_cards
		Foreign-key constraints:
				"fk_users_credit_card" FOREIGN KEY (user_id) REFERENCES users(id)
	*/
	db := tt.Db
	db.Migrator().DropTable(&CreditCard{}, &User{})
	db.Debug().AutoMigrate(&CreditCard{}, &User{})

	/*
		INSERT INTO "users" ("username") VALUES ('Alex3') RETURNING "id"
		INSERT INTO "credit_cards" ("number","user_id") VALUES ('20',1) ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"
	*/
	db.Debug().Create(&User{
		Username:   "Alex3",
		ID:         3,
		CreditCard: CreditCard{Number: "20"},
	})

	// Preload Orders when find users
	/*
		SELECT * FROM "users" WHERE "users"."username" = 'Alex3'
		SELECT * FROM "credit_cards" WHERE "credit_cards"."user_id" = 3
	*/
	users := User{}
	tt.Db.Debug().Preload("CreditCard").Where(&User{Username: "Alex3"}).Find(&users)
	t.Logf("%#v\n", users)

	// db.Migrator().DropTable(&CreditCard{}, &User{})
}
