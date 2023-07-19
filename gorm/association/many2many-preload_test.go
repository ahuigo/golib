package main

import (
	"testing"
	"tt"
)

type Stu struct {
	ID          int `gorm:"primarykey"`
	Stuname     string
	Memberships []Membership `gorm:"foreignKey:StuID"`
}
type School struct {
	ID          int `gorm:"primarykey"`
	Name        string
	Memberships []Membership `gorm:"foreignKey:SchoolID"`
}

type Membership struct {
	ID       int    `gorm:"primarykey"`
	StuID    uint   `gorm:"index;not null"`
	SchoolID uint   `gorm:"index;not null"`
	Stu      Stu    `gorm:"foreignKey:StuID"`
	School   School `gorm:"foreignKey:SchoolID"`
}

func TestMany2Many3rd(t *testing.T) {
	db := tt.Db
	db.Migrator().DropTable(&School{}, "stus", &Membership{})
	db.Debug().AutoMigrate(&Stu{}, &School{}, &Membership{})

	db.Debug().Create(&Membership{
		ID: 3,
		School: School{
			Name: "PKU",
		},
		Stu: Stu{
			Stuname: "Alex3",
		},
	})

	school := School{}
	// 找出PKU 所有的学生
	tt.Db.Debug().Preload("Memberships.Stu").Where("name = ?", "PKU").Find(&school)

	t.Logf("%#v\n", school)

	// db.Migrator().DropTable(&School{}, &Stu{})
}
