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

	record := &BagRelationsTable{
		Model: gorm.Model{ID: 1},
		Data:  &BagRelations{"a": "old"},
	}
	// 1. onconflict时，只能用DoUpdates 手动更新data(因为updateAll 忽略excluded.data，　因为它有default)
	// 1.2 `return　id,data`会修改传入数据的id,data字段(如果没有excluded,就用旧值)
	if err := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"id": gorm.Expr("excluded.id"),
		}),
		UpdateAll: true, // DoUpdates + UpdateAll 会合并
	}).
		Create(record).Error; err != nil {
		t.Fatal(err)
	}

	// 2.1 当onconflictAll时, 忽略excluded.data，　因为它有default)
	// 2.2 `return　id,data`会修改传入数据的id,data字段(如果没有excluded,就用旧值)
	record.Data = &BagRelations{"a": "new"}
	data := record.Data // 必须备份
	err := db.Clauses(clause.OnConflict{
		// Columns:   []clause.Column{{Name: "requirement_id"}, {Name: "task_type"}},
		UpdateAll: true, // DoUpdates + UpdateAll 会合并
	}).Create(&record).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	fmt.Println("data.Data=", record.Data)

	// 3. update 手动更新json字段: Update("data", data2)
	err = tt.Db.Debug().Model(&record).Where("id", record.ID).
		Update("data", data).Error
	if err != nil {
		t.Fatal(err)
	}

	// 4. query
	res := []BagRelationsTable{}
	err = tt.Db.Debug().Find(&res).Error
	if err != nil {
		t.Errorf("err:%v", err)
	}
	fmt.Printf("res:%v\n", res)
}
