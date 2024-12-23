package curd

import (
	"testing"
	"tt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestCreateConflictDoNothing(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	p := []Product{{Code: "L1217", Price: 17, Model: gorm.Model{ID: 1}}}
	tt.Db.Debug().
		Clauses(clause.OnConflict{DoNothing: true}). // 批量时，不冲突的会更新
		Create(&p)
}
func TestCreateConflictUpdateReturn(t *testing.T) {
	type Product1 struct {
		gorm.Model
		Code  string `gorm:"uniqueIndex"`
		Price uint
	}

	tt.Db.AutoMigrate(&Product1{})

	p := &Product1{Code: "L1217", Price: 17}
	err := tt.Db.Debug().
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "code"}}, // 主键、唯一键
			DoUpdates: clause.Assignments(map[string]any{
				// "id": gorm.Expr("excluded.id"), // id会自增 (默认值有语法问题: 少了excluded.)
				"id": gorm.Expr("product1.id"), // id 不会自增(正常不应该自增)
				// "code": gorm.Expr("excluded.code"), // id 不会自增
			}),
		}).
		Create(p).Error
	if err != nil {
		t.Log(err)
	}
	t.Log(p)
}

func TestCreateUpdateAll(t *testing.T) {
	tt.Db.AutoMigrate(&User{})
	db := tt.Db
	users := []User{
		{Username: "usr1", Age: 1, Score: 1},
		{Username: "usr2", Age: 2},
	}
	db.Debug().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}}, //指定联合主键, 默认是primary key(仅限单主键)
			UpdateAll: true,
		}).
		Create(&users)
		/*
			问题1：update All 会更新0值吗？ 会的
		*/
	users = []User{{Username: "usr1", Age: 0, Score: 2}}
	db.Debug().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}}, //指定联合主键, 默认是primary key(仅限单主键)
			UpdateAll: true,
		}).
		Create(&users)
}

// gorm1 移除了： Set("gorm:insert_option", "ON CONFLICT (domain) DO UPDATE SET host= excluded.host,\"remark\"=excluded.remark").
func TestCreateUpdatePartial(t *testing.T) {
	tt.Db.Debug().AutoMigrate(&City{})
	db := tt.Db
	// Update columns to default value on `id` conflict
	users := []User{
		{Age: 1, Score: 1},
	}
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"id":  gorm.Expr("user.id"), // id 不会自增(正常不应该自增)
			"Age": 10,
		}),
		// UpdateAll: true,
	}).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL
}
