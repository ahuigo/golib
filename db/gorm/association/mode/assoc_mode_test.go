package main

import (
	"testing"
	"tt"
)

// doc: https://gorm.io/zh_CN/docs/associations.html
// Association : 用于Join On 查询：你可以使用 Association 来添加、删除、替换、查询join on 关联对象。

func TestAssocModeFind(t *testing.T) {
	db := tt.Db
	models := []interface{}{
		&User{}, &Language{}, &Email{}, &Address{},
	}
	db.Exec("drop table IF EXISTS user_languages;")
	db.Migrator().DropTable(models...)
	db.AutoMigrate(models...)
	db.AutoMigrate(&Email{}, &Address{})
	user := User{
		Name:           "ahuigo",
		BillingAddress: Address{Address1: "Billing Address - Address 2"},
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
	// SELECT "languages"."id","languages"."name" FROM "languages" JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" AND "user_languages"."user_id" = 1
	db.Model(&user).Debug().Association("Languages").Find(&languages)
	// SELECT "languages"."id","languages"."name" FROM "languages" JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" AND "user_languages"."user_id" = 1 WHERE name IN ('zh-CN','en-US')
	db.Model(&user).Debug().Where("name IN ?", codes).Association("Languages").Find(&languages)

	// Append association: 添加一个新的 Email 到 User
	db.Model(&user).Association("Languages").Append(&languages)
	// Replace association:
	db.Model(&user).Association("Languages").Replace(&languages)

	// Count with conditions
	db.Model(&user).Debug().Where("name IN ?", codes).Association("Languages").Count()
	// SELECT count(*) FROM languages JOIN user_languages ON user_languages."language_id" = languages."id" AND "user_languages"."user_id" = 1 WHERE name IN ('zh-CN','en-US')

}
