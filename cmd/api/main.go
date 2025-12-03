package api

import (
	iRoutes "team99_user_service/api/route"
	"team99_user_service/config"
	"team99_user_service/module/database"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	e := echo.New()
	cfg := config.LoadConfig()
	e.Logger.Info("Configuration Loaded")

	db, err := database.InitDB(cfg)
	if err != nil {
		e.Logger.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(); err != nil {
		e.Logger.Fatalf("Failed to run migrations: %v", err)
	}
	e.Logger.Info("Database migrations completed")

	iRoutes.SetupRoutes(e, db, cfg)

	port := ":11231"
	e.Logger.Infof("Listening on port %s", port)
	e.Logger.Fatal(e.Start(port))
}
