package api

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(handler *Handler) *echo.Echo {
	e := echo.New()

	// Группа для /people
	people := e.Group("/people")
	people.Use(LoggingMiddleware)

	people.GET("", handler.FindByLastName)
	people.POST("", handler.CreatePerson)

	return e
}
