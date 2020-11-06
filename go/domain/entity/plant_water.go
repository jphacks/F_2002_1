package entity

import "time"

type PlantWater struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	PlantID   int
	Plant     Plant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WaterID   int
	Water     Water `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	threshold int
}

type PlantWaters []Water
