package request

type CultivationsGetByID struct {
	CultivationID int `param:"cultivation_id"`
}

type CultivationsPutByID struct {
	CultivationID int    `param:"cultivation_id"`
	NickName      string `json:"nick_name"`
}

type CultivationsDeleteByID struct {
	CultivationID int `param:"cultivation_id"`
}
