package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())
	
	e.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "API is ready.")
	})

	// models routes
	ModelsRoutes(e.Group("/models"))
	return e
}