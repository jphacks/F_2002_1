package database

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jphacks/F_2002_1/go/config"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/labstack/echo/v4"

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

func ResetDB(c echo.Context) error {
	givenPass := c.Request().Header.Get("Authorization")
	log.Println(givenPass)
	requiredPass := os.Getenv("ADMIN_PASSW0RD")
	log.Println(requiredPass)
	if givenPass != requiredPass {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	db, err := gorm.Open("mysql", config.DSN())
	if err != nil {
		_ = fmt.Errorf("failed to open MySQL: %w", err)
		return c.String(http.StatusInternalServerError, "Reset Failed")
	}
	db.LogMode(true)
	dropAll(db)
	migrate(db)

	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping(); err != nil {
		_ = fmt.Errorf("failed to ping: %w", err)
		return c.String(http.StatusInternalServerError, "Reset Failed")
	}
	return c.String(http.StatusOK, "Reset Completed!")
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

func dropAll(db *gorm.DB) {
	log.Println("Start drop")
	db.DropTable(&entity.Cultivation{})
	db.DropTable(&entity.Harvesting{})
	db.DropTable(&entity.PlantTemperature{})
	db.DropTable(&entity.PlantWater{})
	db.DropTable(&entity.Plant{})
	db.DropTable(&entity.Season{})
	db.DropTable(&entity.Temperature{})
	db.DropTable(&entity.User{})
	db.DropTable(&entity.Water{})
	db.DropTable(&entity.Watering{})
	log.Println("Finish drop")
}
