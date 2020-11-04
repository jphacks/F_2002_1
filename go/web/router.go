package web

import (
	"log"
	"net/http"

	"github.com/jphacks/F_2002_1/go/database"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewServer はREST APIエンドポイントのハンドラやミドルウェアが登録されたechoの構造体を返却します。
func NewServer(logger *log.Logger) *echo.Echo {
	db, err := database.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	userUC := usecase.NewUserUseCase(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userHandler := handler.NewUserHandler(userUC)

	v1 := e.Group("") // v1 := e.Group("/api/v1")
	v1.GET("/", hello)
	// v1.GET("/admin/reset", ResetDB)

	v1.GET("/users", userHandler.GetUsers)
	v1.GET("/users/:id", userHandler.GetUser)
	v1.POST("/users", userHandler.PostUser)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.DELETE("/users/:id", userHandler.DeleteUser)

	return e
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
