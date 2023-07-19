package main

import (
	"testing"
	"tt"
)

func TestMany2Many(t *testing.T) {
	type School struct {
		ID   int `gorm:"primarykey"`
		Name string
		// Stus []*Stu `gorm:"many2many:stu_schools;"` // 这里注释掉(就不能preload实现通过School 查stus)
	}
	type Stu struct {
		ID      int `gorm:"primarykey"`
		Stuname string
		Schools []School `gorm:"many2many:stu_schools;"`
	}
	/*
		它会创建关联表、CONSTRAINT foreign key:
			ALTER TABLE "stu_schools" ADD CONSTRAINT "fk_stu_schools_stu" FOREIGN KEY ("stu_id") REFERENCES "stus"("id")
			ahuigo=# \d stu_schools
					stu_id     | bigint |           | not null |
					school_id | bigint |           | not null |
			Indexes:
				"stu_schools_pkey" PRIMARY KEY, btree (stu_id, school_id)
	*/
	db := tt.Db
	db.Migrator().DropTable(&School{}, &Stu{}, "stu_schools")
	db.Debug().AutoMigrate(&Stu{})

	/*
		INSERT INTO "stus" ("stuname","id") VALUES ('Alex3',3) RETURNING "id"
		INSERT INTO "schools" ("name") VALUES ('PKU'),('TSU') ON CONFLICT DO NOTHING RETURNING "id"
		INSERT INTO "stu_schools" ("stu_id","school_id") VALUES (3,1),(3,2) ON CONFLICT DO NOTHING
	*/
	db.Debug().Create(&Stu{
		Stuname: "Alex3",
		ID:      3,
		Schools: []School{
			{Name: "PKU"}, {Name: "TSU"},
		},
	})

	// Preload
	/*
		SELECT * FROM "stus" WHERE "stus"."stuname" = 'Alex3'
		SELECT * FROM "stu_schools" WHERE "stu_schools"."stu_id" = 3
		SELECT * FROM "schools" WHERE "schools"."id" IN (1,2)
	*/
	stus := Stu{}
	tt.Db.Debug().Preload("Schools").Where(&Stu{Stuname: "Alex3"}).Find(&stus)
	t.Logf("%#v\n", stus)

	// db.Migrator().DropTable(&School{}, &Stu{})
}
