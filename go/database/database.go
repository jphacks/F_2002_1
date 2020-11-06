package database

import (
	"fmt"
	"log"

	"github.com/jphacks/F_2002_1/go/config"
	"github.com/jphacks/F_2002_1/go/domain/entity"

	"github.com/jinzhu/gorm"
)

// NewDB はMySQLへ接続し、gormのマッピングオブジェクトを生成します。
func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}
	db.LogMode(true)
	migrate(db)

	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return db, nil
}

func migrate(db *gorm.DB) {
	log.Println("Start auto migration")
	db.AutoMigrate(
		&entity.Cultivation{},
		&entity.Harvesting{},
		&entity.PlantTemperature{},
		&entity.PlantWater{},
		&entity.Plant{},
		&entity.Season{},
		&entity.Temperature{},
		&entity.User{},
		&entity.Water{},
		&entity.Watering{},
	)
	log.Println("Finish auto migration")

	log.Println("Start inserting data")
	for _, plant := range Plants {
		db.Create(&plant)
	}
	log.Println("Finish inserting data")
}
