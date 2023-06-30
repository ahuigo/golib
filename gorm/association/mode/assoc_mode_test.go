package main

import (
	"testing"
	"tt"
)

// To skip the auto save when creating/updating, you can use Select or Omit, for example:

func TestAssocModeFind(t *testing.T) {
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

	db.Create(&user)

	codes := []string{"zh-CN", "en-US"}

	// Find matched associations
	languages := []Language{}
	db.Model(&user).Debug().Association("Languages").Find(&languages)
	db.Model(&user).Where("name IN ?", codes).Association("Languages").Find(&languages)

	// Count with conditions
	db.Model(&user).Debug().Where("name IN ?", codes).Association("Languages").Count()
	// SELECT count(*) FROM languages JOIN user_languages ON user_languages."language_id" = languages."id" AND "user_languages"."user_id" = 1 WHERE name IN ('zh-CN','en-US')

}
