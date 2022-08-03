package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

// a successful case
func TestShouldUpdateStats(t *testing.T) {
	gormDb, db, mock := initMockGormDb(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err := recordStats(gormDb, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a failing test case
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	gormDb, db, mock := initMockGormDb(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// expect error
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err := recordStats(gormDb, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func recordStats(db *gorm.DB, userID, productID int64) (err error) {
	tx := db.Begin()

	for {
		if db = tx.Exec("UPDATE products SET views = views + 1"); db.Error != nil {
			break
		}
		if db = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); db.Error != nil {
			break
		}
		break
	}
	err = db.Error
	switch db.Error {
	case nil:
		if db = tx.Commit(); db.Error != nil {
			return db.Error
		}
	default:
		tx.Rollback()
	}
	return err
}
