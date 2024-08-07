package main

import (
	"testing"
	"tt"
)

type UserModel struct {
	// index
	Username string `gorm:"not null;default:'';type:varchar(100)" `
}

func TestMigrateModel(t *testing.T) {
	db := tt.Db
	db.Migrator().DropTable(&UserModel{})
	err := db.Debug().AutoMigrate(
		&UserModel{},
	)
	if err != nil {
		panic(err)
	}
}
