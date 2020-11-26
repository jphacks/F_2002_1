package entity

import "time"

// 一旦不要
// type record struct {
// }

// type records []record

type Cultivation struct {
	ID                  int `gorm:"primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `sql:"index"`
	UserID              int
	User                User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlantID             int
	Plant               Plant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Psoc6DeviceID       int
	StartCultivatingAt  *time.Time
	FinishCultivatingAt *time.Time
	NickName            string
}

type Cultivations []Cultivation
