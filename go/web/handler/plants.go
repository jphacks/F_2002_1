package handler

import (
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

// PlantHandler は /plants 以下のエンドポイントを管理する構造体です。
type PlantHandler struct {
	plantUC *usecase.PlantUseCase
}

// NewPlantHandler はPlantHandlerのポインタを生成する関数です。
func NewPlantHandler(db *gorm.DB) *PlantHandler {
	return &PlantHandler{plantUC: usecase.NewPlantUseCase(db)}
}

// GetPlants は GET /plants に対応するハンドラです。
func (h *PlantHandler) GetPlants(c echo.Context) error {
	logger := log.New()

	plants, err := h.plantUC.ReadPlants()
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := response.PlantGet{}
	for _, plant := range *plants {
		res = append(res, openapi.Plant{
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
		})
	}
	return c.JSON(http.StatusOK, res)
}

// GetPlant は GET /plants/:id に対応するハンドラです。
func (h *PlantHandler) GetPlant(c echo.Context) error {
	logger := log.New()

	req := &request.PlantGetById{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	plant, err := h.plantUC.ReadPlant(req.PlantID)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	res := response.PlantPost{
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
	}
	return c.JSON(http.StatusOK, res)
}

// PostPlants は Post /plants に対応するハンドラです。
func (h *PlantHandler) PostPlants(c echo.Context) error {
	logger := log.New()

	req := &request.PlantPost{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	plant := &entity.Plant{
		Name: req.Name,
	}
	plant, err := h.plantUC.CreatePlant(plant)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	res := response.PlantPost{
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
	}
	return c.JSON(http.StatusOK, res)
}

// UpdatePlant は GET /plants/:id に対応するハンドラです。
func (h *PlantHandler) UpdatePlant(c echo.Context) error {
	logger := log.New()

	req := &request.PlantPost{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	plant := &entity.Plant{
		Name: req.Name,
	}
	plant, err := h.plantUC.UpdatePlant(plant)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := response.PlantPutById{
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
	}
	return c.JSON(http.StatusOK, res)
}

// DeletePlant は Delete /plants/:id に対応するハンドラです。
func (h *PlantHandler) DeletePlant(c echo.Context) error {
	logger := log.New()

	req := &request.PlantDeleteById{}
	if err := c.Bind(req); err != nil {
		logger.Errorj(map[string]interface{}{"message": "failed to bind", "error": err.Error()})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err := h.plantUC.DeletePlant(req.PlantID)
	if err != nil {
		logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, nil)
}
