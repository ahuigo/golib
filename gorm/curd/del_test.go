package curd

import (
	"testing"
	"tt"

	"github.com/samber/lo"
)

func TestDel(t *testing.T) {
	// 自动迁移模式
	tt.Db.AutoMigrate(&Person{})

	p := Person{Name: "com", Age: 3}

	// 指定字段
	tt.Db.Debug().Unscoped().Where("Username=?", "Alex").Delete(p)
	tt.Db.Debug().Unscoped().Delete(p, "Username=? and age=?", "Alex", 23)

	// 指定结构体
	tt.Db.Debug().Unscoped().Where(p).Delete(Person{})
	tt.Db.Debug().Unscoped().Delete(Person{}, p)

}

func TestDelBatch(t *testing.T) {
	type Person struct {
		Name     string `gorm:"primary_key" json:"name" form:"name"`
		Username string `gorm:"unique_index:idx_username" json:"username"`
		Age      int
	}
	tt.Db.AutoMigrate(&Product{})
	ps := []Person{
		{Name: "user0", Age: 3},
		{Name: "user01", Age: 3},
		{Username: "user1"},
		{Username: "user2"},
	}
	tt.Db.Debug().Unscoped().Where(ps).Delete(Person{}) // and condition(wrong)
	tt.Db.Debug().Unscoped().Delete(ps)                 // id in () // Note:没有主键，删除全部

	// batch delete
	delNames := lo.Map(ps, func(p Person, index int) string {
		return p.Username
	})
	tt.Db.Debug().Unscoped().Delete(Person{}, "username in ?", delNames) // and condition
}
