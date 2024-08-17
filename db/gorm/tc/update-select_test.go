package tc

import (
	"testing"
	"tt"
	"tt/curd"
)

func TestInsertOnColumn(t *testing.T) {
	db := tt.Db
	//db.AutoMigrate(&Stock{})
	/*
		// update and select (原子)
		WITH cte AS (
			SELECT id FROM stocks WHERE code='L21' LIMIT 1
		)
		update stocks set code='L21' where id in (select id from cte) RETURNING *;
	*/
	sql := `update stocks set code='L21' where code='L21' RETURNING *`
	results := []curd.Stock{}
	if err := db.Raw(sql).Scan(&results).Error; err != nil {
		t.Error(err)
	} else if len(results) == 0 {
		t.Error("No results")
	} else {
		t.Log("Results:", results)
	}
}
