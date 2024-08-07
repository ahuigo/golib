package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initMockGormDb(t *testing.T) (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	// sqldb
	db, mock, _ := sqlmock.New()
	// defer db.Close()

	// gorm db
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	return gormDb, db, mock
}
