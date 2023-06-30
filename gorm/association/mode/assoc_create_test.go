package main

import (
	"testing"
	"tt"

	"gorm.io/gorm"
)

type Address struct {
	UID      uint `gorm:"primarykey"`
	Address1 string
}
type Email struct {
	ID    uint `gorm:"primarykey"`
	UID   uint
	Email string
}

// User has and belongs to many languages, `user_languages` is the join table
type User struct {
	ID               uint `gorm:"primarykey"`
	Name             string
	BillingAddressID uint
	BillingAddress   Address    `gorm:"foreignKey:uid;references:id;"`
	Emails           []Email    `gorm:"foreignKey:uid;references:id;"`
	Languages        []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Users []*User `gorm:"many2many:user_languages;"` // 如果不用preload 这里可省略
}

func TestAssoication(t *testing.T) {
	db := tt.Db
	models := []interface{}{
		&User{}, &Language{}, &Email{}, &Address{},
	}
	db.Debug().Migrator().DropTable(models...)
	db.Debug().AutoMigrate(models...)
	db.Debug().AutoMigrate(&Email{}, &Address{})
	user := User{
		Name:           "ahuigo",
		BillingAddress: Address{Address1: "Billing Address - Address 1"},
		Emails: []Email{
			{Email: "jinzhu@example.com"},
			{Email: "jinzhu-2@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	// db.Omit(clause.Associations).Create(&user)
	db.Debug().Create(&user)

	t.Logf("-------------------------------------------------")
	// BEGIN TRANSACTION;
	// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
	// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
	// COMMIT;

	db.Debug().Save(&user) // upsert
	t.Logf("----------session---------------------------------------")
	db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

}
