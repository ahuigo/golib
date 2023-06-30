package main

import (
	"testing"
	"tt"
)

func TestMany2Many(t *testing.T) {
	type Language struct {
		ID   int `gorm:"primarykey"`
		Name string
		// Users []*User `gorm:"many2many:user_languages;"`
	}
	type User struct {
		ID        int `gorm:"primarykey"`
		Username  string
		Languages []Language `gorm:"many2many:user_languages;"`
	}
	/*
		它会创建关联表、CONSTRAINT foreign key:
			ALTER TABLE "user_languages" ADD CONSTRAINT "fk_user_languages_user" FOREIGN KEY ("user_id") REFERENCES "users"("id")
			ahuigo=# \d user_languages
					user_id     | bigint |           | not null |
					language_id | bigint |           | not null |
			Indexes:
				"user_languages_pkey" PRIMARY KEY, btree (user_id, language_id)
	*/
	db := tt.Db
	db.Migrator().DropTable(&Language{}, &User{})
	db.Debug().AutoMigrate(&User{})

	/*
		INSERT INTO "users" ("username","id") VALUES ('Alex3',3) RETURNING "id"
		INSERT INTO "languages" ("name") VALUES ('English'),('Chinese') ON CONFLICT DO NOTHING RETURNING "id"
		INSERT INTO "user_languages" ("user_id","language_id") VALUES (3,1),(3,2) ON CONFLICT DO NOTHING
	*/
	db.Debug().Create(&User{
		Username: "Alex3",
		ID:       3,
		Languages: []Language{
			{Name: "English"}, {Name: "Chinese"},
		},
	})

	// Preload
	/*
		SELECT * FROM "users" WHERE "users"."username" = 'Alex3'
		SELECT * FROM "user_languages" WHERE "user_languages"."user_id" = 3
		SELECT * FROM "languages" WHERE "languages"."id" IN (1,2)
	*/
	users := User{}
	tt.Db.Debug().Preload("Languages").Where(&User{Username: "Alex3"}).Find(&users)
	t.Logf("%#v\n", users)

	// db.Migrator().DropTable(&Language{}, &User{})
}
