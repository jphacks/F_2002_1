package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web/handler/openapi"
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
