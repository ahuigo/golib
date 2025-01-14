package tt

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var SqlDb *sql.DB

func GetDb() *gorm.DB {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	_ = newLogger
	// schema prefix: https://stackoverflow.com/questions/54203807/how-set-a-specifc-database-schema-with-gorm
	dsn := "host=localhost user=role1 password='' dbname=ahuigo port=5432 sslmode=disable search_path=public TimeZone=Asia/Shanghai"
	//dsn:="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: newLogger,
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// Db = Db.Debug()

	SqlDb, _ = Db.DB()
	// defer sqlDB.Close()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDb.SetConnMaxLifetime(time.Hour)

	//Debug
	if os.Getenv("DEBUG") == "true" {
		Db = Db.Debug()
	}

	return Db
}
func init() {
	GetDb()
}
