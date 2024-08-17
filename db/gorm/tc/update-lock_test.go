package tc

import (
	"runtime"
	"sync"
	"testing"
	"time"
	"tt"
	"tt/curd"

	"gorm.io/gorm/clause"
)

/*
* refer: post/db/pg/pg-lock.md
 */
func TestUpdateLock(t *testing.T) {
	var err error
	var wg sync.WaitGroup
	tt.Db.Migrator().DropTable(&curd.Person{})
	tt.Db.AutoMigrate(&curd.Person{})

	// 1. insert
	p := curd.Person{Name: "com", Username: "Alex", Age: 3, Addrs: []string{"a", "b"}}
	if err = tt.Db.Debug().Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// 2. tx
	wg.Add(2)
	go func() {
		defer wg.Done()
		res := curd.Person{}
		tx := tt.Db.Debug().Begin()
		defer tx.Commit()

		// 3. select for update
		// tx.Set("gorm:query_option", "FOR UPDATE SKIP LOCKED")
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).Where("username=?", p.Username).Find(&res).Error; err != nil {
			panic(err)
		} else {
			if len(res.Addrs) == 0 {
				// panic("res.Addrs is empty")
				t.Fatal("res.Addrs is empty")
			}
		}
		time.Sleep(5 * time.Second)
	}()

	// 4. update
	go func() {
		defer wg.Done()
		t.Log("update...")
		if err := tt.Db.Debug().Model(&curd.Person{}).Where("username=?", p.Username).Update("age", 4).Error; err != nil {
			t.Log(err)
			runtime.Goexit()
		}
		t.Log("update done")
	}()
	wg.Wait()
	println("done")

}
