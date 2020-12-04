package controller

import (
	"strconv"

	"github.com/jphacks/F_2002_1/go/entity/model"
	"github.com/jphacks/F_2002_1/go/interface/database"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/labstack/echo/v4"
)

type ExampleController struct {
	Interactor usecase.ExampleInteractor
}

func NewExampleController(sqlHandler database.SqlHandler) *ExampleController {
	return &ExampleController{
		Interactor: usecase.ExampleInteractor{
			ExampleRepository: &database.ExampleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ExampleController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	example, err := controller.Interactor.ExampleById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, example)
	return
}

func (controller *ExampleController) Index(c echo.Context) (err error) {
	cities, err := controller.Interactor.Examples()
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, cities)
	return
}

func (controller *ExampleController) Create(c echo.Context) (err error) {
	ex := model.Example{}
	_ = c.Bind(&ex)
	example, err := controller.Interactor.Add(ex)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, example)
	return
}

func (controller *ExampleController) Save(c echo.Context) (err error) {
	ex := model.Example{}
	_ = c.Bind(&ex)
	example, err := controller.Interactor.Update(ex)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, example)
	return
}

func (controller *ExampleController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	example := model.Example{
		ID: id,
	}
	err = controller.Interactor.DeleteById(example)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
