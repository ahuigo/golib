package main

import (
	"testing"
	"tt"
)

// doc: https://gorm.io/docs/many_to_many.html#Override-Foreign-Key
type Company struct {
	ID   int    `gorm:"primarykey"`
	Name string `gorm:"index:,unique"`
	Stus []Stu  `gorm:"many2many:stu_companys;foreignKey:Name;joinForeignKey:CompanyName;References:Stuname;joinReferences:Stuname;"`
}
type Stu struct {
	ID        int       `gorm:"primarykey"`
	Stuname   string    `gorm:"index:,unique"`
	Companies []Company `gorm:"many2many:stu_companys;"`
}

// type StuCompany struct {
// 	CompanyName string `gorm:"primaryKey"`
// 	Stuname    string `gorm:"primaryKey"`
// }

func TestMany2ManyUniqueKey(t *testing.T) {
	db := tt.Db
	db.Debug().Migrator().DropTable(&Company{}, &Stu{}, "stu_companys")
	db.Debug().AutoMigrate(&Stu{})

	db.Debug().Create(&Stu{
		Stuname: "Alex3",
		ID:      3,
		Companies: []Company{
			{Name: "PKU1"}, {Name: "TSU2"},
		},
	})
	if err := db.Model(&Stu{Stuname: "Alex3"}).Association("Companies").Append(&[]Company{
		{Name: "PKU"}, {Name: "TSU"},
	}); err != nil {
		panic(err)
	}

	stus := Stu{}
	tt.Db.Debug().Preload("Companies").Where(&Stu{Stuname: "Alex3"}).Find(&stus)
	t.Logf("%#v\n", stus)
	// tt.Db.Debug().Preload("Stus").Where(&Stu{Stuname: "Alex3"}).Find(&stus)

	company := Company{Name: "TSU"}
	tt.Db.Debug().Preload("Stus").Where(&company).Find(&company)
	t.Logf("%#v\n", company)
}
