package entity

import "time"

type Harvesting struct {
	ID            int `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	CultivationID int
	Cultivation   Cultivation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Harvestings []Harvesting
