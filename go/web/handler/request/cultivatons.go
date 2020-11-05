package request

type CultivationsGet struct {
	CultivationID int `param:"cultivation_id"`
}

type CultivationsPut struct {
	CultivationID int    `param:"cultivation_id"`
	NickName      string `json:"nick_name"`
}

type CultivationsDelete struct {
	CultivationID int `param:"cultivation_id"`
}
