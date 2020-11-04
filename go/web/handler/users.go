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

// UsersHandler は /users 以下のエンドポイントを管理する構造体です。
type UsersHandler struct {
	userUC *usecase.UserUseCase
}

// NewUsersHandler はUsersHandlerのポインタを生成する関数です。
func NewUsersHandler(db *gorm.DB) *UsersHandler {
	return &UsersHandler{userUC: usecase.NewUserUseCase(db)}
}

// GetUsers は GET /users に対応するハンドラです。
func (h *UsersHandler) GetUsers(c echo.Context) error {
	logger := log.New()
	ctx := c.Request().Context()

	users, err := h.userUC.ReadUsers(ctx)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser は GET /users/:id に対応するハンドラです。
func (h *UsersHandler) GetUser(c echo.Context) error {
	logger := log.New()
	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.userUC.ReadUser(ctx, id)
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

// PostUser は POST /users に対応するハンドラです。
func (h *UsersHandler) PostUser(c echo.Context) error {
	logger := log.New()
	user := new(entity.User) // req
	if err := c.Bind(user); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ctx := c.Request().Context()
	user, err := h.userUC.CreateUser(ctx, user)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser は PUT /users/:id に対応するハンドラです。
func (h *UsersHandler) UpdateUser(c echo.Context) error {
	logger := log.New()
	user := new(entity.User) // req
	if err := c.Bind(user); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.userUC.UpdateUser(ctx, id, user)
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

// DeleteUser は DELETE /users/:id に対応するハンドラです。
func (h *UsersHandler) DeleteUser(c echo.Context) error {
	logger := log.New()
	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.userUC.DeleteUser(ctx, id)
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
