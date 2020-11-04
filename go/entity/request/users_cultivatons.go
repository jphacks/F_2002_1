package request

type UsersCultivationsGet struct {
	UserId int `json:"user_id"`
}

type UsersCultivationsPost struct {
	UserId   int    `json:"user_id"`
	PlantId  int    `json:"plant_id"`
	NickName string `json:"nick_name"`
}
