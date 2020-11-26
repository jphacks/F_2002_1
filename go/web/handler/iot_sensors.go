package handler

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web/handler/request"
	"github.com/jphacks/F_2002_1/go/web/handler/response"
	"github.com/labstack/echo/v4"
)

type IotSensorsHandler struct {
	cultivationUC *usecase.CultivationUseCase
	userUC        *usecase.UserUseCase
}

func NewIotSenorsHandler(db *gorm.DB) *IotSensorsHandler {
	return &IotSensorsHandler{cultivationUC: usecase.NewCultivationUseCase(db), userUC: usecase.NewUserUseCase(db)}
}

func (h *IotSensorsHandler) PostIotSensors(c echo.Context) error {
	logger := log.New()

	req := &request.IotSensorsPost{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user, err := h.userUC.ReadUser(req.UserID)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			logger.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if req.RaspiDeviceID != user.RaspiDeviceID {
		logger.Printf("UserID: %d does not have RaspiDeviceID: %d", req.UserID, req.RaspiDeviceID)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	for _, data := range req.Data {
		cultivation, err := h.cultivationUC.ReadCultivation(data.CultivationID)
		if err != nil {
			if errors.Is(err, entity.ErrCultivationNotFound) {
				logger.Debug(err)
				return echo.NewHTTPError(http.StatusNotFound)
			}
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if data.Psoc6DeviceID != cultivation.Psoc6DeviceID {
			logger.Printf("CultivationID: %d does not have Psoc6DeviceID: %d", data.CultivationID, data.Psoc6DeviceID)
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		//TODO: チャット送信フラグ追加
		humidityStatus := ""
		temperatureStatus := ""
		illuminanceStatus := ""
		soliMoistureStatus := ""
		pressureStatus := ""

		_ = &response.IotSensorsPost{
			CultivationID: data.CultivationID,
			CreatedAt:     data.CreatedAt,
			Humidity: response.IotSensorData{
				Value:  data.Humidity,
				Status: humidityStatus,
			},
			Temperature: response.IotSensorData{
				Value:  data.Temperature,
				Status: temperatureStatus,
			},
			Illuminance: response.IotSensorData{
				Value:  data.Illuminance,
				Status: illuminanceStatus,
			},
			SoilMoisture: response.IotSensorData{
				Value:  data.SoilMoisture,
				Status: soliMoistureStatus,
			},
			Pressure: response.IotSensorData{
				Value:  data.Pressure,
				Status: pressureStatus,
			},
		}
	}

	//TODO: firebaseに送信

	return c.JSON(http.StatusCreated, nil)
}
