package curd

import (
	"testing"
	"tt"

	"github.com/lib/pq"
)

func TestWhereArray(t *testing.T) {
	// Where("entity_ids && ?", pq.StringArray(cellIds)
	p := Stock{Code: "L1217", Price: 19}
	_ = tt.Db.Debug().Model(&p).Where("entity_ids && ?", pq.StringArray([]string{"1"})).Updates(&p).Error
}
