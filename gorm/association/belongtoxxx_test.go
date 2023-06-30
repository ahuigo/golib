package main

import (
	"database/sql"
	"testing"
	"tt"

	"gorm.io/gorm"
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

func TestBelongtoX(t *testing.T) {
	db := tt.Db
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

	company2 := BelongChild{}
	db.Model(&user).Preload("BelongParent").Where(&BelongChild{ParentID: company.Pid}).Find(&company2)
	t.Logf("%#v\n", user)

}
