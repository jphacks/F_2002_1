package entity

import "time"

type PlantPressure struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	PlantID    int
	Plant      Plant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status     string
	LowerValue float32
}

type PlantPressures []PlantPressure
