package handler

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"

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
func (h *CultivationsHandler) GetCultivations(c echo.Context) error {
	logger := log.New()
	uid := c.Param("id")

	cultivation, err := h.cultivationUC.ReadCultivationsByCultivationId(uid)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, cultivation)
}

// UpdateCultivation は PUT /cultivations/:id に対応するハンドラです。
func (h *CultivationsHandler) UpdateCultivation(c echo.Context) error {
	logger := log.New()
	cultivation := new(entity.Cultivation) // req
	if err := c.Bind(cultivation); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	id := c.Param("id")

	cultivation, err := h.cultivationUC.UpdateCultivation(id, cultivation)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, cultivation)
}

// DeleteCultivation は DELETE /cultivation に対応するハンドラです。
func (h *CultivationHandler) DeleteCultivation(c echo.Context) error {
	logger := log.New()
	uid := c.Request().Header.Get("Authorization")

	cultivation, err := h.cultivationUC.DeleteCultivationByUid(uid)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, cultivation)
}
