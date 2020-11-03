package request

type PlantGet struct{}

type PlantGetById struct {
	PlantId int `json:"plant_id"`
}
