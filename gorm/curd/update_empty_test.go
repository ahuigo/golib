package curd

import (
	"testing"
	"tt"
)

func TestUpdateEmpty(t *testing.T) {
	// 自动迁移模式
	var err error
	tt.Db.Migrator().DropTable(&Person{})
	tt.Db.AutoMigrate(&Person{})

	// 1. insert
	p := Person{Name: "com", Username: "Alex", Age: 3, Addrs: []string{"a", "b"}}
	if err = tt.Db.Debug().Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2. do not update empty
	p2 := Person{Name: "com", Username: "Alex2", Age: 0, Addrs: []string(nil)}
	// UPDATE "people" SET "name"='com',"username"='Alex2' WHERE name='com'
	if err := tt.Db.Model(&Person{}).Debug().Where("name=?", p.Name).Updates(p2).Error; err != nil {
		t.Fatal(err)
	}
	// 3. update empty via Select('*') and Omit("addrs")
	// UPDATE "people" SET "name"='com',"username"='Alex2',"age"=0,"valid"=NULL WHERE name='com'
	if err := tt.Db.Model(&Person{}).Debug().Select("*").Omit("addrs").Where("name=?", p.Name).Updates(p2).Error; err != nil {
		t.Fatal(err)
	}

}
