package api

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/xjncx/people-info-api/docs"
)

func NewRouter(handler *Handler) *echo.Echo {
	e := echo.New()

	people := e.Group("/people")
	people.Use(LoggingMiddleware)

	people.GET("/search", handler.FindByLastName)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
