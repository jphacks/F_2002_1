package main

import (
	"os"

	"github.com/jphacks/F_2002_1/go/config"
	"github.com/jphacks/F_2002_1/go/database"
	"github.com/jphacks/F_2002_1/go/log"
	"github.com/jphacks/F_2002_1/go/usecase"
	"github.com/jphacks/F_2002_1/go/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := log.New()

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

	userRepo := database.NewUserRepository(db)

	userUC := usecase.NewUserUseCase(userRepo)

	s := web.NewServer(userUC)

	logger.Print(config.Port())
	if err := s.Start(":" + config.Port()); err != nil {
		logger.Infof("shutting down the server with error: %v", err)
		os.Exit(1)
	}
}