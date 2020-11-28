package response

import "github.com/jphacks/F_2002_1/go/web/handler/openapi"

type UserCultivationsGetById openapi.Cultivation

// type UserCultivationsPost openapi.Cultivation
type UserCultivationsPost struct {
	ID int `json:"id"`
}

type UserCultivationsPutById openapi.Cultivation

type UserCultivationsDeleteById struct{}
