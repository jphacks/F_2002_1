package openapi

import "time"

type Plant struct {
	ID           int           `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
	Name         string        `json:"name"`
	NickName     string        `json:"nick_name"`
	Price        int           `json:"price"`
	Period       int           `json:"period"`
	Difficulty   int           `json:"difficulty"`
	Description  string        `json:"description"`
	KitName      string        `json:"kit_name"`
	Season       Season        `json:"season"`
	Temperatures []Temperature `json:"temperatures"`
	Waters       []Water       `json:"waters"`
}
