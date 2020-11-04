package request

type CultivationsGet struct {
	CultivationId int `json:"cultivation_id"`
}

type CultivationsPut struct {
	CultivationId int    `json:"cultivation_id"`
	NickName      string `json:"nick_name"`
}

type CultivationsDelete struct {
	CultivationId int `json:"cultivation_id"`
}
