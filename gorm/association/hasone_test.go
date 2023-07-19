package main

import (
	"database/sql"
	"testing"
	"tt"

	"gorm.io/gorm"
)

// `BelongParent` belongs to `BelongChild`, `ChildID` is the foreign key(依赖child)
type BelongParent struct {
	gorm.Model
	Name         string
	Age          uint8
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	ChildID      uint
	BelongChild  BelongChild `gorm:"foreignKey:child_id;references:pid;"` // BelongParent(child_id) has one child
	// BelongChild   BelongChild `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BelongChild struct {
	Pid   uint `gorm:"primarykey"`
	Cname string
}

/*
*
常用于group: belong to company
*/
func TestHasOne(t *testing.T) {
	db := tt.Db
	child := BelongChild{
		Pid:   1,
		Cname: "child1",
	}
	parent := BelongParent{
		ChildID: child.Pid,
		Name:    "parent1",
	}
	db.Migrator().DropTable(&BelongChild{}, &BelongParent{})
	db.AutoMigrate(&BelongChild{}, &BelongParent{})
	db.Create(&child)
	db.Create(&parent)

	parent2 := BelongParent{}
	db.Model(&parent2).Preload("BelongChild").Where(&BelongParent{ChildID: child.Pid}).Find(&parent2)
	t.Logf("%#v\n", parent2)

}
