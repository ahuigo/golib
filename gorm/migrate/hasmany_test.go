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

// HasmanyParent has one HasmanyChild, HasmanyChildID is the foreign key(parentID)
type HasmanyParent struct {
	ID              uint           `gorm:"primarykey"`
	HasmanyChildren []HasmanyChild `gorm:"foreignKey:HasmanyParentID;references:id;"`
	// HasmanyChild HasmanyChild
}

type HasmanyChild struct {
	ID              uint `gorm:"primarykey"`
	Name            string
	HasmanyParentID uint
}

func TestHasmany(t *testing.T) {
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

	db.Migrator().DropTable(&HasmanyChild{}, &HasmanyParent{})
	db.AutoMigrate(&HasmanyChild{}, &HasmanyParent{})

	user := &HasmanyParent{
		HasmanyChildren: []HasmanyChild{
			{Name: "card1"},
			{Name: "card2"},
		},
	}
	db.Create(user)
	// INSERT INTO "hasmany_parents" DEFAULT VALUES RETURNING "id"
	// INSERT INTO "hasmany_children" ("name","hasmany_parent_id") VALUES('card1',1),('card2',1)
	// ON CONFLICT ("id") DO UPDATE SET "hasmany_parent_id"="excluded"."hasmany_parent_id" RETURNING "id"

}
