package main

import (
	"testing"
	"tt"
)

// doc: https://gorm.io/docs/many_to_many.html#Override-Foreign-Key
// usage: https://github.com/go-gorm/gorm/issues/6482
// 自动生成中间表：StuCompany
/**
companies		->		stu_companys					-->						stus
Name(foreignKey)		CompanyName(joinForeignKey)/Stuname(joinReferences)		Stuname(References)

*/
type Company struct {
	ID   int    `gorm:"primarykey"`
	Name string `gorm:"index:,unique"`
	Stus []Stu  `gorm:"many2many:stu_companys;foreignKey:Name;joinForeignKey:CompanyName;joinReferences:Stuname;References:Stuname;"`
}

/*
*
companies		<--		stu_companys						<--					stus
Name(References)		CompanyName(joinReferences)/Stuname(joinForeignKey)		Stuname(foreignKey)
*/
type Stu struct {
	ID      int    `gorm:"primarykey"`
	Stuname string `gorm:"index:,unique"`
	// Companies []Company `gorm:"many2many:stu_companys;"` // default join key is "stu_id",ERROR: column "stu_id" of "stu_companys" does not exist
	Companies []Company `gorm:"many2many:stu_companys;foreignKey:Stuname;joinForeignKey:Stuname;joinReferences:CompanyName;References:Name;"`
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

	stu := Stu{}
	tt.Db.Debug().Preload("Companies").Where(&Stu{Stuname: "Alex3"}).Find(&stu)
	t.Logf("stu:%#v\n", stu)
	// tt.Db.Debug().Preload("Stus").Where(&Stu{Stuname: "Alex3"}).Find(&stus)

	company := Company{Name: "TSU"}
	tt.Db.Debug().Preload("Stus").Where(&company).Find(&company)
	t.Logf("company:%#v\n", company)
}
