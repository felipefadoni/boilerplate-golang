package postgres

import (
	"os"
	"time"

	"github.com/felipefadoni/boilerplate-golang/src/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetInstance() *gorm.DB {
	db := connectDb()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(100)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1000)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func InitDatabase() {
	db := connectDb()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.Query(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	db.AutoMigrate(entities.User{})

}
