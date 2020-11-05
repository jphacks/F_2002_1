package request

type PlantGet struct{}

type PlantGetById struct {
	PlantID int `param:"plant_id"`
}

type PlantPost struct {
	Name string `json:"name"`
}

type PlantPutById struct {
	PlantID int    `param:"plant_id"`
	Name    string `json:"name"`
}

type PlantDeleteById struct {
	PlantID int `param:"plant_id"`
}
