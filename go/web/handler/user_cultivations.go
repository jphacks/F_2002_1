package handler

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web/fbauth"
	"github.com/jphacks/F_2002_1/go/web/handler/openapi"
	"github.com/jphacks/F_2002_1/go/web/handler/request"
	"github.com/jphacks/F_2002_1/go/web/handler/response"

	"github.com/labstack/echo/v4"
)

// CultivationsHandler は /cultivations 以下のエンドポイントを管理する構造体です。
type UserCultivationsHandler struct {
	cultivationUC *usecase.CultivationUseCase
	userUC        *usecase.UserUseCase
	plantUC       *usecase.PlantUseCase
}

// NewCultivationsHandler はCultivationsHandlerのポインタを生成する関数です。
func NewUserCultivationsHandler(db *gorm.DB) *UserCultivationsHandler {
	return &UserCultivationsHandler{cultivationUC: usecase.NewCultivationUseCase(db)}
}

// GetUser は GET /cultivations/:id に対応するハンドラです。
func (h *UserCultivationsHandler) GetUserCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.UserCultivationsGetByID{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
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

	isExist, err := h.cultivationUC.CheckCultivationByIDUID(req.CultivationID, uid)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if !isExist {
		return echo.NewHTTPError(http.StatusForbidden)
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

	plant, err := h.plantUC.ReadPlant(cultivation.PlantID)
	if err != nil {
		if errors.Is(err, entity.ErrPlantNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := &response.UserCultivationsGetById{
		ID:        cultivation.ID,
		CreatedAt: cultivation.CreatedAt,
		UpdatedAt: cultivation.UpdatedAt,
		DeletedAt: cultivation.DeletedAt,
		Plant: openapi.Plant{
			ID:          plant.ID,
			CreatedAt:   plant.CreatedAt,
			UpdatedAt:   plant.UpdatedAt,
			DeletedAt:   plant.DeletedAt,
			Name:        plant.Name,
			NickName:    plant.NickName,
			Price:       plant.Price,
			Period:      plant.Period,
			Difficulty:  plant.Difficulty,
			Description: plant.Description,
			KitName:     plant.KitName,
			Season: openapi.Season{
				From: plant.SeasonFrom,
				To:   plant.SeasonTo,
			},
		},
		NickName:            cultivation.NickName,
		StartCultivatingAt:  *cultivation.StartCultivatingAt,
		FinishCultivatingAt: *cultivation.FinishCultivatingAt,
	}

	return c.JSON(http.StatusOK, res)
}

// PostUser は POST /cultivations に対応するハンドラです。
func (h *UserCultivationsHandler) PostUserCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.UserCultivationsPost{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
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

	cultivation := &entity.Cultivation{UserID: uid} // TODO
	cultivation, err = h.cultivationUC.CreateCultivation(cultivation)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	plant, err := h.plantUC.ReadPlant(cultivation.PlantID)
	if err != nil {
		if errors.Is(err, entity.ErrPlantNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := &response.UserCultivationsPost{
		ID:        cultivation.ID,
		CreatedAt: cultivation.CreatedAt,
		UpdatedAt: cultivation.UpdatedAt,
		DeletedAt: cultivation.DeletedAt,
		Plant: openapi.Plant{
			ID:          plant.ID,
			CreatedAt:   plant.CreatedAt,
			UpdatedAt:   plant.UpdatedAt,
			DeletedAt:   plant.DeletedAt,
			Name:        plant.Name,
			NickName:    plant.NickName,
			Price:       plant.Price,
			Period:      plant.Period,
			Difficulty:  plant.Difficulty,
			Description: plant.Description,
			KitName:     plant.KitName,
			Season: openapi.Season{
				From: plant.SeasonFrom,
				To:   plant.SeasonTo,
			},
		},
		NickName:            cultivation.NickName,
		StartCultivatingAt:  *cultivation.StartCultivatingAt,
		FinishCultivatingAt: *cultivation.FinishCultivatingAt,
	}

	return c.JSON(http.StatusCreated, res)
}

// UpdateCultivation は PUT /cultivations/:id に対応するハンドラです。
func (h *UserCultivationsHandler) UpdateUserCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.UserCultivationsPutByID{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
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

	isExist, err := h.cultivationUC.CheckCultivationByIDUID(req.CultivationID, uid)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if !isExist {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	cultivation := &entity.Cultivation{ID: req.CultivationID, UserID: uid} // TODO
	cultivation, err = h.cultivationUC.UpdateCultivation(cultivation)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	plant, err := h.plantUC.ReadPlant(cultivation.PlantID)
	if err != nil {
		if errors.Is(err, entity.ErrPlantNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := &response.UserCultivationsPutById{
		ID:        cultivation.ID,
		CreatedAt: cultivation.CreatedAt,
		UpdatedAt: cultivation.UpdatedAt,
		DeletedAt: cultivation.DeletedAt,
		Plant: openapi.Plant{
			ID:          plant.ID,
			CreatedAt:   plant.CreatedAt,
			UpdatedAt:   plant.UpdatedAt,
			DeletedAt:   plant.DeletedAt,
			Name:        plant.Name,
			NickName:    plant.NickName,
			Price:       plant.Price,
			Period:      plant.Period,
			Difficulty:  plant.Difficulty,
			Description: plant.Description,
			KitName:     plant.KitName,
			Season: openapi.Season{
				From: plant.SeasonFrom,
				To:   plant.SeasonTo,
			},
		},
		NickName:            cultivation.NickName,
		StartCultivatingAt:  *cultivation.StartCultivatingAt,
		FinishCultivatingAt: *cultivation.FinishCultivatingAt,
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteCultivation は DELETE /cultivation に対応するハンドラです。
func (h *UserCultivationsHandler) DeleteUserCultivation(c echo.Context) error {
	logger := log.New()

	req := &request.UserCultivationsDeleteByID{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
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

	isExist, err := h.cultivationUC.CheckCultivationByIDUID(req.CultivationID, uid)
	if err != nil {
		if errors.Is(err, entity.ErrCultivationNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if !isExist {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	err = h.cultivationUC.DeleteCultivation(req.CultivationID)
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
