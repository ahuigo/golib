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

// M2user has and belongs to many languages, `m2user_languages` is the join table
type M2user struct {
	gorm.Model
	Languages []Language `gorm:"many2many:m2user_languages;"`
}

type Language struct {
	gorm.Model
	Name    string
	M2users []*M2user `gorm:"many2many:m2user_languages;"` // 如果不用preload 这里可省略
}
type M2userLanguage struct{}

func TestMany2many(t *testing.T) {
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

	db.Migrator().DropTable(&M2user{}, &Language{}, &M2userLanguage{})
	db.AutoMigrate(&M2user{}, &Language{})
	/******
		CREATE TABLE "m2user_languages" (
	     "language_id" BIGINT,
	     "m2user_id"   BIGINT,
	     PRIMARY KEY ("language_id", "m2user_id"),
	     CONSTRAINT "fk_m2user_languages_language" FOREIGN KEY ("language_id")
			REFERENCES "languages"("id"),
	     CONSTRAINT "fk_m2user_languages_m2user" FOREIGN KEY ("m2user_id")
			REFERENCES "m2users"("id")
	  )
	*/

}
