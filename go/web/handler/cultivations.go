package handler

import (
	"errors"
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
type CultivationsHandler struct {
	cultivationUC *usecase.CultivationUseCase
}

// NewCultivationsHandler はCultivationsHandlerのポインタを生成する関数です。
func NewCultivationsHandler(db *gorm.DB) *CultivationsHandler {
	return &CultivationsHandler{cultivationUC: usecase.NewCultivationUseCase(db)}
}

// GetCultivation は GET /cultivations/:id に対応するハンドラです。
func (h *CultivationsHandler) GetCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.CultivationsGetByID{}
	if err := c.Bind(&req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation, err := h.cultivationUC.ReadCultivation(req.CultivationID)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
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
		NickName:            cultivation.NickName,
		StartCultivatingAt:  cultivation.StartCultivatingAt,
		FinishCultivatingAt: cultivation.FinishCultivatingAt,
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateCultivation は PUT /cultivations/:id に対応するハンドラです。
func (h *CultivationsHandler) UpdateCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.CultivationsPutByID{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation := &entity.Cultivation{
		ID:       req.CultivationID,
		NickName: req.NickName,
	}
	cultivation, err := h.cultivationUC.UpdateCultivation(cultivation)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
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
		NickName:            cultivation.NickName,
		StartCultivatingAt:  cultivation.StartCultivatingAt,
		FinishCultivatingAt: cultivation.FinishCultivatingAt,
	}
	return c.JSON(http.StatusOK, res)
}

// DeleteCultivation は DELETE /cultivation に対応するハンドラです。
func (h *CultivationsHandler) DeleteCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.CultivationsDeleteByID{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err := h.cultivationUC.DeleteCultivation(req.CultivationID)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusNoContent, nil)
}
