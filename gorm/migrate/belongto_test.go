package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// `BelongChild` belongs to `BelongParent`, `ParentID` is the foreign key(依赖parent)
type BelongChild struct {
	gorm.Model
	Name         string
	Age          uint8
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	ParentID     uint
	BelongParent BelongParent `gorm:"foreignKey:parent_id;references:pid;"`
	// BelongParent   BelongParent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}

type BelongParent struct {
	Pid   uint `gorm:"primarykey"`
	Cname string
}

func TestBelongto(t *testing.T) {
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

	company := BelongParent{
		Pid: 1,
	}
	user := BelongChild{
		ParentID: company.Pid,
	}
	db.Migrator().DropTable(&BelongParent{}, &BelongChild{})
	db.AutoMigrate(&BelongParent{}, &BelongChild{})
	db.Create(&company)
	db.Create(&user)

	db.First(&user, 1)
	db.Model(&user).Update("Age", 18)
	db.Model(&user).Omit("Role").Updates(map[string]interface{}{"Name": "jinzhu", "Role": "admin"})
	db.Unscoped().Delete(&user)

}
