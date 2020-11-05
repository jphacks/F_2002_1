package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"

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

// GetCultivationsByUID は GET /users/:id/cultivations に対応するハンドラです。
func (h *UsersCultivationsHandler) GetUsersCultivationsByUID(c echo.Context) error {
	logger := log.New()

	uid, _ := strconv.Atoi(c.Param("id"))
	cultivation, err := h.cultivationUC.ReadCultivationsByUID(uid)
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

// PostUsersCultivation は POST /user/cultivations に対応するハンドラです。
func (h *UsersCultivationsHandler) PostUsersCultivation(c echo.Context) error {
	logger := log.New()

	uid, _ := strconv.Atoi(c.Param("id"))
	cultivation := &entity.Cultivation{UserID: uid} // req
	if err := c.Bind(cultivation); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation, err := h.cultivationUC.CreateCultivation(cultivation)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, cultivation)
}
