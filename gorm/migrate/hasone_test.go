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

// HasoneParent has one HasoneChild, HasoneChildID is the foreign key(child.parentID)
type HasoneParent struct {
	ID          uint        `gorm:"primarykey"`
	HasoneChild HasoneChild `gorm:"foreignKey:HasoneParentID;references:id;"`
	// HasoneChild HasoneChild
}

type HasoneChild struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	HasoneParentID uint
}

func TestHasone(t *testing.T) {
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

	db.Migrator().DropTable(&HasoneChild{}, &HasoneParent{})
	db.AutoMigrate(&HasoneChild{}, &HasoneParent{})

	user := &HasoneParent{
		HasoneChild: HasoneChild{
			Name: "card1",
		},
	}
	db.Create(user)
	// INSERT INTO "hasone_parents" DEFAULT VALUES RETURNING "id"
	// INSERT INTO "hasone_children" ("name","hasone_parent_id") VALUES ('card1',1) ON CONFLICT ("id") DO UPDATE SET "hasone_parent_id"="excluded"."hasone_parent_id" RETURNING "id"

}
