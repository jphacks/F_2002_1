package openapi

import "time"

type User struct {
	ID           int           `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
	Uid          string        `json:"uid"`
	Name         string        `json:"name"`
	Cultivations []Cultivation `json:"cultivations"`
}
