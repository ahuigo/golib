package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	gorm.Model
	Name             string
	BillingAddressID uint
	BillingAddress   Address    `gorm:"foreignKey:uid;references:id;"`
	Emails           []Email    `gorm:"foreignKey:uid;references:id;"`
	Languages        []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"` // 如果不用preload 这里可省略
}
type UserLanguage struct{}

func TestAssoication(t *testing.T) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	dsn := "host=localhost host=localhost user=role1 dbname=ahuigo password= sslmode=disable TimeZone=Asia/Shanghai"
	dsn = "postgres://role1:@localhost:5432/ahuigo?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	fmt.Println(err)

	models := []interface{}{
		&User{}, &Language{}, &Email{}, &Address{},
	}
	db.Migrator().DropTable(&UserLanguage{})
	db.Migrator().DropTable(models...)
	db.AutoMigrate(models...)
	user := User{
		Name:           "jinzhu",
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
	db.Create(&user)
	// BEGIN TRANSACTION;
	// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
	// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
	// COMMIT;

	// db.Save(&user)

}
