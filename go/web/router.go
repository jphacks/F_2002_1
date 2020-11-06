package web

import (
	"net/http"

	"github.com/jphacks/F_2002_1/go/database"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/web/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewServer はREST APIエンドポイントのハンドラやミドルウェアが登録されたechoの構造体を返却します。
func NewServer() *echo.Echo {
	logger := log.New()

	db, err := database.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	// defer func() {
	// 	err := db.Close()
	// 	if err != nil {
	// 		logger.Fatal(err)
	// 	}
	// }()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1 := e.Group("") // v1 := e.Group("/api/v1")
	v1.GET("/", hello)
	v1.PUT("/admin/reset", database.ResetDB)

	plantsHandler := handler.NewPlantHandler(db)
	v1.GET("/plants", plantsHandler.GetPlants)

	usersHandler := handler.NewUsersHandler(db)
	v1.GET("/users", usersHandler.GetUsers)
	v1.GET("/users/:id", usersHandler.GetUser)
	v1.POST("/users", usersHandler.PostUser)
	v1.PUT("/users/:id", usersHandler.UpdateUser)
	v1.DELETE("/users/:id", usersHandler.DeleteUser)

	userHandler := handler.NewUserHandler(db)
	v1.GET("/user", userHandler.GetUser)
	v1.PUT("/user", userHandler.UpdateUser)
	v1.DELETE("/user", userHandler.DeleteUser)

	usersCultivationsHandler := handler.NewUsersCultivationsHandler(db)
	v1.POST("/users/:id/cultivations", usersCultivationsHandler.PostUsersCultivation)

	cultivationsHandler := handler.NewCultivationsHandler(db)
	v1.GET("/cultivations/:id", cultivationsHandler.GetCultivation)
	v1.PUT("/cultivations/:id", cultivationsHandler.UpdateCultivation)
	v1.DELETE("/cultivations/:id", cultivationsHandler.DeleteCultivation)

	userCultivationsHandler := handler.NewUserCultivationsHandler(db)
	v1.POST("/user/cultivations", userCultivationsHandler.PostUserCultivation)
	v1.GET("/user/cultivations/:id", userCultivationsHandler.GetUserCultivation)
	v1.PUT("/user/cultivations/:id", userCultivationsHandler.UpdateUserCultivation)
	v1.DELETE("/user/cultivations/:id", userCultivationsHandler.DeleteUserCultivation)

	return e
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
