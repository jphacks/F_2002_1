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

// UserHandler は /users 以下のエンドポイントを管理する構造体です。
type UserHandler struct {
	userUC *usecase.UserUseCase
}

// NewUserHandler はUserHandlerのポインタを生成する関数です。
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{userUC: usecase.NewUserUseCase(db)}
}

// GetUser は GET /user に対応するハンドラです。
func (h *UserHandler) GetUser(c echo.Context) error {
	logger := log.New()

	req := &request.UserGet{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uid, err := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		if errors.Is(err, entity.ErrInvalidIdToken) {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	id, err := h.userUC.ReadUserIDByUID(uid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user, err := h.userUC.ReadUser(id)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	resCultivations := []openapi.Cultivation{}
	for _, cultivation := range user.Cultivations {
		resCultivations = append(resCultivations, openapi.Cultivation{
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
		})
	}
	res := &response.UserGet{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
		Uid:          user.Uid,
		Name:         user.Name,
		Cultivations: resCultivations,
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateUser は PUT /user に対応するハンドラです。
func (h *UserHandler) UpdateUser(c echo.Context) error {
	logger := log.New()

	req := &request.UserPut{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uid, err := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		if errors.Is(err, entity.ErrInvalidIdToken) {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	id, err := h.userUC.ReadUserIDByUID(uid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user := &entity.User{
		ID:   id,
		Name: req.Name,
	}
	user, err = h.userUC.UpdateUser(user)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	resCultivations := []openapi.Cultivation{}
	for _, cultivation := range user.Cultivations {
		resCultivations = append(resCultivations, openapi.Cultivation{
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
		})
	}
	res := &response.UserGet{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
		Uid:          user.Uid,
		Name:         user.Name,
		Cultivations: resCultivations,
	}
	return c.JSON(http.StatusOK, res)
}

// DeleteUser は DELETE /user に対応するハンドラです。
func (h *UserHandler) DeleteUser(c echo.Context) error {
	logger := log.New()

	req := &request.UserDelete{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uid, err := fbauth.GetUIDByToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		if errors.Is(err, entity.ErrInvalidIdToken) {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	id, err := h.userUC.ReadUserIDByUID(uid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = h.userUC.DeleteUser(id)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusNoContent, nil)
}
