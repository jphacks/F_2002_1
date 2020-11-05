package request

type UsersCultivationsGet struct {
	UserID int `param:"user_id"`
}

type UsersCultivationsPost struct {
	UserID   int    `param:"user_id"`
	PlantID  int    `json:"plant_id"`
	NickName string `json:"nick_name"`
}
