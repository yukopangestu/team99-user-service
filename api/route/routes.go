package route

import (
	authmiddleware "team99_listing_service/api/middleware"
	"team99_listing_service/config"
	Ihandler "team99_listing_service/module/handler"
	Irepository "team99_listing_service/module/repository"
	Iservice "team99_listing_service/module/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1 := e.Group("/api/v1")
	{
		authMiddleware := authmiddleware.NewAuthMiddleware(cfg.JWTServiceKey)
		userRepository := Irepository.NewUserRepository(db)
		userService := Iservice.NewUserService(userRepository)
		userHandler := Ihandler.NewUserHandler(userService)

		internalListing := v1.Group("listings")
		internalListing.Use(authMiddleware.ValidateSecretKey())
		internalListing.GET("/", userHandler.GetListing)
		internalListing.POST("/", userHandler.CreateListing)

		//This API doesn't need different handler because it has the same request and response with the internal one
		publicListing := v1.Group("public-api/listings")
		publicListing.GET("/", userHandler.GetListing)
		publicListing.POST("/", userHandler.CreateListing)
	}
}
