package request

type CultivationsGetByID struct {
	CultivationID int `param:"id"`
}

type CultivationsPutByID struct {
	CultivationID int    `param:"id"`
	NickName      string `json:"nick_name"`
}

type CultivationsDeleteByID struct {
	CultivationID int `param:"id"`
}
