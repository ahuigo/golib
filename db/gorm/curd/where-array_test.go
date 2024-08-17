package curd

import (
	"testing"
	"tt"

	"github.com/lib/pq"
)

func TestWherePlainArray(t *testing.T) {
	dest := []Person{}
	// error:  SELECT * FROM "people" WHERE "people"."addrs" = '{"a2","b2"}'
	p := Person{Addrs: pq.StringArray([]string{"a2", "b2"})}
	tt.Db.Debug().Model(&Person{}).Find(&dest, &p)

	//  SELECT * FROM "people"
	p = Person{Addrs: nil}
	tt.Db.Debug().Model(&Person{}).Find(&dest, &p)
}

func TestWhereElemInQueryArray(t *testing.T) {
	p := Stock{Code: "L1217", Price: 19}
	//  WHERE code = any('{"L21"}') AND "code" = 'L1217'
	_ = tt.Db.Debug().Model(&p).Where("code = any(?)", pq.StringArray([]string{"L21"})).Updates(&p).Error
	// WHERE code in ('L21') AND "code" = 'L1217'
	_ = tt.Db.Debug().Model(&p).Where("code in (?)", []string{"L21"}).Updates(&p).Error
}

func TestWhereArrayAny(t *testing.T) {
	dest := []Person{}
	tt.Db.Debug().Model(&Person{}).Where("'addr1' = ANY(addrs)").Find(&dest)

}

func TestWhereArrayIntersectArray(t *testing.T) {
	// SELECT * FROM "people" WHERE addrs && '{"a2","b2"}'
	p := Person{Addrs: pq.StringArray([]string{"a2", "b2"})}
	dest := []Person{}
	tt.Db.Debug().Model(&Person{}).Where("addrs && ?", p.Addrs).Find(&dest)
}
func TestWhereArrayIncludeArray(t *testing.T) {
	// SELECT * FROM "people" WHERE addrs @> '{"a2","b2"}'
	p := Person{Addrs: pq.StringArray([]string{"a2", "b2"})}
	dest := []Person{}
	tt.Db.Debug().Model(&Person{}).Where("addrs @> ?", p.Addrs).Find(&dest)
}

func TestWhereArrayEqualStrict(t *testing.T) {
	// SELECT * FROM "people" WHERE "people"."addrs" = '{"a2","b2"}';
	p := Person{Addrs: pq.StringArray([]string{"a2", "b2"})}
	dest := []Person{}
	tt.Db.Debug().Model(&Person{}).Find(&dest, &p)
}
