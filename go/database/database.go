package database

import (
	"fmt"

	"github.com/jphacks/F_2002_1/go/config"

	"github.com/jinzhu/gorm"
)

// NewDB はMySQLへ接続し、gormのマッピングオブジェクトを生成します。
func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}
	db.LogMode(true)

	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return db, nil
}
