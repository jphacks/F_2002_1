package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/web/handler/request"
	"github.com/labstack/echo/v4"
)

type IotSensorsHandler struct {
	// cultivationUC s*usecase.CultivationUseCase
	// userUC        *usecase.UserUseCase
}

func NewIotSensorsHandler(db *gorm.DB) *IotSensorsHandler {
	// return &IoTSensorsHandler{cultivationUC: usecase.NewCultivationUseCase(db), userUC: usecase.NewUserUseCase(db)}
	return &IotSensorsHandler{}
}

func (h *IotSensorsHandler) PostIotSensors(c echo.Context) error {
	logger := log.New()

	req := &request.IotSensorsPost{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	//TODO: 認証
	//TODO: チャット送信フラグ追加
	//TODO: firebaseに送信

	return c.JSON(http.StatusCreated, nil)
}
