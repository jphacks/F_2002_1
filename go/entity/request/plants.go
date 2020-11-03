package request

type PlantGet struct{}

type PlantGetById struct {
	PlantId int `json:"plant_id"`
}

type PlantPost struct {
	Name string `json:"name"`
}

type PlantPutById struct {
	PlantId int    `json:"plant_id"`
	Name    string `json:"name"`
}

type PlantDeleteById struct {
	PlantId int `json:"plant_id"`
}
