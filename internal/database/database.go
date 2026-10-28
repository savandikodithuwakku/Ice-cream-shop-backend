package database

import (
	"fmt"
	"time"

	"icecream-backend/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// Apply foreign key enforcement to every database connection.
	dsn := "icecream.db?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get database connection: %w", err)
	}

	// Keep the initial SQLite setup simple.
	sqlDB.SetMaxOpenConns(1)

	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Item{},
		&models.Customer{},
	)
	if err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("create tables: %w", err)
	}

	return db, nil
}
