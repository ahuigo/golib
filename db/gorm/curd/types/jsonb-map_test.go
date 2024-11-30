package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"tt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BagRelations map[string]string

// Value makes `BagRelations` satisfy the `driver.Valuer` interface.
func (j BagRelations) Value() (driver.Value, error) {
	value, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// 很重要: 用于指定pg的数据类型
func (j BagRelations) GormDataType() string {
	// return "BagRelations"
	return "jsonb"
}

// Scan makes `BagRelations` satisfy the `sql.Scanner` interface.
func (j *BagRelations) Scan(value interface{}) error {
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

// DefaultValueInterface (没有用，不接受自定义有类型)
func (u *BagRelations) DefaultValueInterface() map[string]interface{} {
	return map[string]interface{}{
		"name": "alex",
	}
}

type BagRelationsTable struct {
	gorm.Model
	Data *BagRelations `json:"data" gorm:"not null;default:'{}'"`
}

func TestMyBagRelations(t *testing.T) {
	db := tt.Db.Debug()
	if err := tt.Db.AutoMigrate(&BagRelationsTable{}); err != nil {
		t.Fatal(err)
	}

	datas := []BagRelationsTable{
		{
			Data: &BagRelations{"a": "b"},
		},
	}

	// 2. upsert(Note:无法insert更新json(因为包括default), 且会覆盖data的值)
	data := datas[0]
	data.Data = &BagRelations{"a": "300", "b": "4"}
	data2 := data.Data // 必须备份
	err := db.Clauses(clause.OnConflict{
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
	res := []BagRelationsTable{}
	err = tt.Db.Debug().Find(&res).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	fmt.Printf("res:%v\n", res)
}
