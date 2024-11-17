package types

import (
	"fmt"
	"strings"
	"testing"
	"tt"
)

func TestCrudBytea(t *testing.T) {
	type KV struct {
		Mykey string `gorm:"primaryKey;type:varchar(2048);not null"`
		Myval []byte `gorm:"type:bytea"`
	}
	p := KV{Mykey: "k2", Myval: []byte("v19")}
	db := tt.GetDb()
	db.Migrator().AutoMigrate(&KV{})
	err := db.Create(&p).Error
	if err != nil && !strings.Contains(err.Error(), "duplicate key value") {
		t.Fatalf("err:%v\n", err)
	}
	// db.Create
	cols := [][]byte{}
	err = db.Debug().Model(&p).Select("myval").Pluck("myval", &cols).Error
	fmt.Printf("val:%v, err:%v\n", cols, err)

}
