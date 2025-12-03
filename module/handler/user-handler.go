package handler

import (
	"net/http"
	"team99_listing_service/module/model"
	"team99_listing_service/module/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// In handler struct or as global
var validate = validator.New()

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(listingService service.UserServiceInterface) *UserHandler {
	return &UserHandler{service: listingService}
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	var id string = c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"result":  false,
			"message": "Invalid request body",
		})
	}

	data, err := h.service.GetUserById(id)
	if err != nil {
		c.Logger().Error("Failed to get listing, got error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"result":  false,
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result":   true,
		"listings": data,
	})
}

func (h *UserHandler) CreateListing(c echo.Context) error {
	var request model.PostListingRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"result":  false,
			"message": "Invalid request body",
		})
	}

	if err := validate.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"result":  false,
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	data, err := h.service.PostListing(request)
	if err != nil {
		c.Logger().Error("Failed to create listing, got error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"result":  false,
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result":  true,
		"listing": data,
	})
}
