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
	uid := c.Request().Header.Get("Authorization")

	user, err := h.userUC.ReadUserByUid(uid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser は PUT /user に対応するハンドラです。
func (h *UserHandler) UpdateUser(c echo.Context) error {
	logger := log.New()
	user := new(entity.User) // req
	if err := c.Bind(user); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uid := c.Request().Header.Get("Authorization")

	user, err := h.userUC.UpdateUserByUid(uid, user)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser は DELETE /user に対応するハンドラです。
func (h *UserHandler) DeleteUser(c echo.Context) error {
	logger := log.New()
	uid := uid := c.Request().Header.Get("Authorization")

	user, err := h.userUC.DeleteUserByUid(uid)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}
