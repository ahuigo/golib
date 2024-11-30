package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"tt"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// 很重要
func (j JSONB) GormDataType() string {
	return "jsonb"
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
	Data  *JSONB       `json:"data" gorm:"not null;type:jsonb;default:'{}'"`
	Data2 pgtype.JSONB `json:"data2" gorm:"type:jsonb;default:'[]';not null"`
}

func TestJsonb(t *testing.T) {
	o := pgtype.JSONB{}
	json.Unmarshal([]byte(`{"a":1}`), &o)
	o2 := struct {
		A pgtype.JSONB `json:"a"`
	}{
		A: o,
	}
	res, err := json.Marshal(o2)
	if err != nil {
		t.Fatal(err)
	}
	println(string(res))
}
func TestMyJsonb(t *testing.T) {
	tt.Db.AutoMigrate(&JsonbTable{})

	datas := []JsonbTable{
		{Data: &JSONB{"a": 1, "b": 2}},
		{Data: &JSONB{"a": 2, "b": 3}},
		{},
	}

	// 1. create
	err := tt.Db.Debug().Create(&datas).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}

	// 2. upsert(Note:无法insert更新json(因为包括default), 且会覆盖data的值)
	data := datas[0]
	data.Data = &JSONB{"a": 300, "b": 4}
	data2 := data.Data // 必须备份
	err = tt.Db.Debug().Clauses(clause.OnConflict{
		// Columns:   []clause.Column{{Name: "requirement_id"}, {Name: "task_type"}},
		UpdateAll: true,
	}).Create(&data).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	// 2.2 update(更新json)
	err = tt.Db.Debug().Model(&data).Where("id", data.ID).
		Update("data", data2).Error
	if err != nil {
		t.Fatal(err)
	}

	// 3. query
	res := []JsonbTable{}
	err = tt.Db.Debug().Find(&res).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	fmt.Printf("res:%v\n", res)
}
