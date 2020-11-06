package request

type UsersCultivationsGet struct {
	UserID int `param:"id"`
}

type UsersCultivationsPost struct {
	UserID   int    `param:"id"`
	PlantID  int    `json:"plant_id"`
	NickName string `json:"nick_name"`
}
