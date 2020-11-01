package infrastracture

import (
	"net/http"

	"github.com/jphacks/F_2002_1/go/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// IniitApiServer is
func InitApiServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	exampleController := controller.NewExampleController(NewSqlHandler())

	e.GET("/", hello)
	e.GET("/admin/reset", ResetDB)

	e.GET("/examples", func(c echo.Context) error { return exampleController.Index(c) })
	e.GET("/examples/:id", func(c echo.Context) error { return exampleController.Show(c) })
	e.POST("/examples", func(c echo.Context) error { return exampleController.Create(c) })
	e.PUT("/examples/:id", func(c echo.Context) error { return exampleController.Save(c) })
	e.DELETE("/examples/:id", func(c echo.Context) error { return exampleController.Delete(c) })

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
