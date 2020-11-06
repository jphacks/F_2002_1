package request

type UserCultivationsGetByID struct {
	CultivationID int `param:"cultivation_id"`
}

type UserCultivationsPost struct {
	PlantID  int    `json:"plant_id"`
	NickName string `json:"nick_name"`
}

type UserCultivationsPutByID struct {
	CultivationID int    `param:"cultivation_id"`
	NickName      string `json:"nick_name"`
}

type UserCultivationsDeleteByID struct {
	CultivationID int `param:"cultivation_id"`
}
