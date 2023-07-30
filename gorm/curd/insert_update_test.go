package curd

import (
	"testing"
	"tt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestCreateConflict(t *testing.T) {
	tt.Db.AutoMigrate(&Product{})

	p := []Product{{Code: "L1217", Price: 17, Model: gorm.Model{ID: 1}}}
	tt.Db.Debug().
		Clauses(clause.OnConflict{DoNothing: true}). // 批量时，不冲突的会更新
		Create(&p)
}

func TestCreateUpdateAll(t *testing.T) {
	tt.Db.AutoMigrate(&User{})
	db := tt.Db
	users := []User{
		{UserName: "usr1", Age: 1},
		{UserName: "usr2", Age: 2},
	}
	db.Debug().
		Clauses(clause.OnConflict{
			// Columns:   []clause.Column{{Name: "groupname","username"}}, // 默认是primary key(仅限单主键，非联合主键)
			UpdateAll: true,
		}).
		Create(&users)
}

// gorm1 移除了： Set("gorm:insert_option", "ON CONFLICT (domain) DO UPDATE SET host= excluded.host,\"remark\"=excluded.remark").
func TestCreateUpdate(t *testing.T) {
	tt.Db.Debug().AutoMigrate(&City{})
	db := tt.Db
	// Update columns to default value on `id` conflict
	users := []User{}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
	}).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL
}
