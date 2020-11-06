package entity

import "time"

type Watering struct {
	ID            int `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	CultivationID int
}

type Waterings []Watering
