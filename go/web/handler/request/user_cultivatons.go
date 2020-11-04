package request

type UserCultivationsGetByID struct {
	IdToken       string `json:"Authorization"`
	CultivationId int    `json:"cultivation_id"`
}

type UserCultivationsPost struct {
	IdToken  string `json:"Authorization"`
	PlantId  int    `json:"plant_id"`
	NickName string `json:"nick_name"`
}

type UserCultivationsPutById struct {
	IdToken       string `json:"Authorization"`
	CultivationId int    `json:"cultivation_id"`
	NickName      string `json:"nick_name"`
}

type UserCultivationsDeleteById struct {
	IdToken       string `json:"Authorization"`
	CultivationId int    `json:"cultivation_id"`
}
