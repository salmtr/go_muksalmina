package controller

import (
	"code_competence/model/payload"
	"net/http"

	"code_competence/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	req := payload.CreateUserRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := usecase.CreateUser(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	payload := payload.LoginRequest{}
	c.Bind(&payload)

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := usecase.LoginUser(payload.Email, payload.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
