package curd

import (
	"testing"
	"tt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey,not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func TestCreate(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})
}
