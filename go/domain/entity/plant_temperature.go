package entity

import "time"

type PlantTemperature struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	PlantID   int
	Plant     Plant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WaterID   int
	Water     Water `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Threshold int
}

type PlantTemperatures []Temperature
