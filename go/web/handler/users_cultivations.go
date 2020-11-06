package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web/handler/openapi"
	"github.com/jphacks/F_2002_1/go/web/handler/request"
	"github.com/jphacks/F_2002_1/go/web/handler/response"

	"github.com/labstack/echo/v4"
)

// CultivationsHandler は /cultivations 以下のエンドポイントを管理する構造体です。
type UsersCultivationsHandler struct {
	cultivationUC *usecase.CultivationUseCase
}

// NewCultivationsHandler はCultivationsHandlerのポインタを生成する関数です。
func NewUsersCultivationsHandler(db *gorm.DB) *UsersCultivationsHandler {
	return &UsersCultivationsHandler{cultivationUC: usecase.NewCultivationUseCase(db)}
}

// PostUsersCultivation は POST /user/cultivations に対応するハンドラです。
func (h *UsersCultivationsHandler) PostUsersCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.UsersCultivationsPost{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation := &entity.Cultivation{UserID: req.UserID} // TODO
	cultivation, err := h.cultivationUC.CreateCultivation(cultivation)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := &response.CultivationsGetById{
		ID:        cultivation.ID,
		CreatedAt: cultivation.CreatedAt,
		UpdatedAt: cultivation.UpdatedAt,
		DeletedAt: cultivation.DeletedAt,
		Plant: openapi.Plant{
			ID:          cultivation.Plant.ID,
			CreatedAt:   cultivation.Plant.CreatedAt,
			UpdatedAt:   cultivation.Plant.UpdatedAt,
			DeletedAt:   cultivation.Plant.DeletedAt,
			Name:        cultivation.Plant.Name,
			NickName:    cultivation.Plant.NickName,
			Price:       cultivation.Plant.Price,
			Period:      cultivation.Plant.Period,
			Difficulty:  cultivation.Plant.Difficulty,
			Description: cultivation.Plant.Description,
			KitName:     cultivation.Plant.KitName,
			Season: openapi.Season{
				From: cultivation.Plant.SeasonFrom,
				To:   cultivation.Plant.SeasonTo,
			},
		},
	}

	return c.JSON(http.StatusCreated, res)
}
