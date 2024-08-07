package hook

import (
	"context"
	"fmt"
	"testing"
	"tt"

	"github.com/ahuigo/glogger"
	"gorm.io/gorm"
)

// var db *gorm.DB

func TestBeforeCreate(t *testing.T) {
	type Stock struct {
		Code  string `gorm:"primary_key" `
		Price uint
		Count *uint `json:"count"  gorm:"default:2"`
	}
	db := tt.Db
	db.AutoMigrate(&Stock{})
	// create hook
	beforeCreate := func(scope *gorm.DB) {
		fmt.Println("before create sql")
		contextScopeKey := "my_key"
		rctx, _ := scope.Get(contextScopeKey)
		ctx, ok := rctx.(context.Context)
		if !ok || ctx == nil {
			ctx = context.Background()
		}
		// handle ctx ....
		scope.Set(contextScopeKey, ctx)
	}
	tt.Db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", beforeCreate)

	// create stock
	createStock := func() {
		p := Stock{Code: "L21", Price: 21}
		// no err here
		if err := db.Create(&p).Error; err != nil {
			glogger.Info(err)
			return
		}
	}
	createStock()
}
