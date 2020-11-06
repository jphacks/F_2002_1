package openapi

import "time"

type record struct {
	Waterings   []Watering   `json:"watering"`
	Harvestings []Harvesting `json:"harvesting"`
}

type Cultivation struct {
	ID                  int        `json:"id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at"`
	Plant               Plant      `json:"plant"`
	NickName            string     `json:"nick_name"`
	StartCultivatingAt  *time.Time `json:"start_cultivating_at"`
	FinishCultivatingAt *time.Time `json:"finish_cultivating_at"`
	Record              record     `json:"record"`
}
