package response

import "time"

type Season struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type Temperature struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type Water struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type Plant struct {
	ID           int           `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
	Name         string        `json:"name"`
	NickName     string        `json:"nick_name"`
	Price        string        `json:"price"`
	Period       int           `json:"period"`
	Difficulty   int           `json:"difficulty"`
	Description  string        `json:"description"`
	KitName      string        `json:"kit_name"`
	Season       Season        `json:"season"`
	Temperatures []Temperature `json:"temperatures"`
	Waters       []Water       `json:"waters"`
}

type PlantGet []Plant

type PlantGetById Plant

type PlantPost Plant

type PlantPutById Plant

type PlantDeleteById struct{}
