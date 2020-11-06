package entity

import "time"

type Plant struct {
	ID                int `gorm:"primary_key"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
	Name              string
	NickName          string
	Price             int
	Period            int
	Difficulty        int
	Description       string
	KitName           string
	SeasonFrom        int
	SeasonTo          int
	PlantTemperatures []PlantTemperature
	PlantWaters       []PlantWater
}

type Plants []Plant
