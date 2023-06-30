package main

import (
	"testing"
	"tt"

	"gorm.io/gorm/clause"
)

// To skip the auto save when creating/updating, you can use Select or Omit, for example:

func TestAssocSkip(t *testing.T) {
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

	db.Debug().Select("Name").Create(&user) // skip
	// INSERT INTO "users" (name) VALUES ("jinzhu", 1, 2);

	t.Log("-----------------------")
	db.Omit("BillingAddress").Create(&user)
	// Skip create BillingAddress when creating a user

	t.Log("-----------------------")
	db.Omit(clause.Associations).Create(&user)
	// Skip all associations when creating a user

}
