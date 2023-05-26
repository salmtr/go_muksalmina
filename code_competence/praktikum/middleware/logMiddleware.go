package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logmiddleware(e *echo.Echo) {
	config := middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${method} ${uri} | ${latency_human}\n",
	}
	e.Use(middleware.LoggerWithConfig(config))

}
