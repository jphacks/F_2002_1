package entity

import "time"

type PlantWater struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	PlantID   int
	WaterID   int
	threshold int
}

type PlantWaters []Water
