package curd

import (
	"context"
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Create from customized data type
type Location struct {
	X, Y int
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (loc Location) GormDataType() string {
	// brew install postgis
	// CREATE EXTENSION postgis;
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}

type City struct {
	Name     string
	Location Location
}

func TestCreateType(t *testing.T) {
	tt.Db.Debug().AutoMigrate(&City{})
	tt.Db.Model(City{}).Create(map[string]interface{}{
		"Name":     "jinzhu",
		"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	})
	// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));
	tt.Db.Debug().Create(&City{
		Name:     "jinzhu",
		Location: Location{X: 100, Y: 100},
	})
}
