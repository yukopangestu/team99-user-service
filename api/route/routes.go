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

		internalUser := v1.Group("users")
		internalUser.Use(authMiddleware.ValidateSecretKey())
		internalUser.GET("/", userHandler.GetAllUser)
		internalUser.GET("/:id", userHandler.GetUserById)
		internalUser.POST("/", userHandler.CreateListing)

		//This API doesn't need different handler because it has the same request and response with the internal one
		publicUser := v1.Group("public-api/users")
		publicUser.POST("/", userHandler.CreateListing)
	}
}
