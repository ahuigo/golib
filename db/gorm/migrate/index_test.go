package main

import (
	"testing"
	"tt"

	"gorm.io/gorm"
)

// primary key
type ProductPrimary struct {
	gorm.Model
	// alter table products add column price serial;
	// 相当于
	// CREATE SEQUENCE mytable_item_id_seq OWNED BY mytable.item_id;
	// ALTER TABLE mytable ALTER item_id SET DEFAULT nextval('mytable_item_id_seq');
	ID        uint           `gorm:"primaryKey,AUTO_INCREMENT,not null"` // primaryKey 默认为 AUTO_INCREMENT
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func TestMigrateProductPrimary(t *testing.T) {
	db := tt.Db
	db.Migrator().DropTable(&ProductPrimary{})
	err := db.Debug().AutoMigrate(
		&ProductPrimary{}, //flag1ZZ
	)
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&ProductPrimary{})
}

// https://gorm.io/docs/indexes.html
type User2 struct {
	// index
	Name string `gorm:"index;not null;default:inited"`

	// unique index
	Name4 string `gorm:"uniqueIndex:idx_board_cell;not null"`
	Name5 string `gorm:"uniqueIndex:idx_board_cell;not null"`

	// custom index
	Name2 string `gorm:"index:idx_name,unique"`
	Name3 string `gorm:"index:,sort:desc,type:btree,length:10,where:name3 != 'jinzhu'"`
	Age   int64  `gorm:"index:,comment:hello \\, world,where:age > 10"`
	Age2  int64  `gorm:"index:,expression:ABS(age)"`
	// "idx_user2_status" btree (status) WHERE status = ANY (ARRAY['active'::text, 'archived'::text])
	Status string `gorm:"index:,where: status in ('active'\\,'archived')"`
}

// Composite Indexes, for example:
// create composite index `idx_member` with columns `name`, `number`
type UserCompositeIndexes struct {
	Pk1 string `gorm:"primaryKey"`
	Pk2 string `gorm:"primaryKey"`

	Name   string `gorm:"index:idx_member"`
	Number string `gorm:"index:idx_member"`
}

func TestUser2Index(t *testing.T) {
	db := tt.Db
	db.Migrator().DropTable(&User2{})
	err := db.Debug().AutoMigrate(
		&User2{},
	)
	if err != nil {
		panic(err)
	}
}

func TestUserCompositeIndexes(t *testing.T) {
	db := tt.Db
	db.Migrator().DropTable(&UserCompositeIndexes{})
	err := db.Debug().AutoMigrate(
		&UserCompositeIndexes{},
	)
	if err != nil {
		panic(err)
	}
}
