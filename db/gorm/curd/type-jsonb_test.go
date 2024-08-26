package curd

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
)

type JSONB map[string]interface{}

// Value makes `JSONB` satisfy the `driver.Valuer` interface.
func (j JSONB) Value() (driver.Value, error) {
	value, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// Scan makes `JSONB` satisfy the `sql.Scanner` interface.
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	err := json.Unmarshal(bytes, &j)
	if err != nil {
		return err
	}
	return nil
}

type JsonbTable struct {
	gorm.Model
	Data *JSONB `json:"data" gorm:"not null;type:jsonb;default:'{}'"`
}

func TestJsonb(t *testing.T) {
	tt.Db.AutoMigrate(&JsonbTable{})

	datas := []JsonbTable{
		{Data: &JSONB{"a": 1, "b": 2}},
		{Data: &JSONB{"a": 2, "b": 3}},
		{},
	}
	err := tt.Db.Debug().Create(&datas).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	res := []JsonbTable{}
	err = tt.Db.Debug().Find(&res).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	fmt.Printf("res:%v\n", res)
}
