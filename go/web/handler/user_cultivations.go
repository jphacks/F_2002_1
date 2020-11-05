package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web/fbauth"

	"github.com/labstack/echo/v4"
)

// CultivationsHandler は /cultivations 以下のエンドポイントを管理する構造体です。
type UserCultivationsHandler struct {
	cultivationUC *usecase.CultivationUseCase
	userUC        *usecase.UserUseCase
}

// NewCultivationsHandler はCultivationsHandlerのポインタを生成する関数です。
func NewUserCultivationsHandler(db *gorm.DB) *UserCultivationsHandler {
	return &UserCultivationsHandler{cultivationUC: usecase.NewCultivationUseCase(db)}
}

// GetUser は GET /cultivations/:id に対応するハンドラです。
func (h *UserCultivationsHandler) GetUserCultivation(c echo.Context) error {
	logger := log.New()

	id, _ := strconv.Atoi(c.Param("id"))
	cultivation, err := h.cultivationUC.ReadCultivation(id)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uuid := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	uid, err := h.userUC.ReadUserIDByUID(uuid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if uid != cultivation.UserID {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	return c.JSON(http.StatusOK, cultivation)
}

// PostUser は POST /cultivations に対応するハンドラです。
func (h *UserCultivationsHandler) PostUserCultivation(c echo.Context) error {
	logger := log.New()

	uuid := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	uid, err := h.userUC.ReadUserIDByUID(uuid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation := &entity.Cultivation{UserID: uid}
	if err := c.Bind(cultivation); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cultivation, err = h.cultivationUC.CreateCultivation(cultivation)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, cultivation)
}

// UpdateCultivation は PUT /cultivations/:id に対応するハンドラです。
func (h *UserCultivationsHandler) UpdateUserCultivation(c echo.Context) error {
	logger := log.New()

	id, _ := strconv.Atoi(c.Param("id"))

	uuid := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	uid, err := h.userUC.ReadUserIDByUID(uuid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	isExist, _ := h.cultivationUC.CheckCultivationByIDUID(id, uid)
	if !isExist {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	cultivation := &entity.Cultivation{ID: id, UserID: uid}
	if err := c.Bind(cultivation); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation, err = h.cultivationUC.UpdateCultivation(cultivation)
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
func (h *UserCultivationsHandler) DeleteUserCultivation(c echo.Context) error {
	logger := log.New()

	id, _ := strconv.Atoi(c.Param("id"))

	uuid := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	uid, err := h.userUC.ReadUserIDByUID(uuid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	isExist, _ := h.cultivationUC.CheckCultivationByIDUID(id, uid)
	if !isExist {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	cultivation := &entity.Cultivation{ID: id, UserID: uid}
	if err := c.Bind(cultivation); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cultivation, err = h.cultivationUC.DeleteCultivation(id)
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
