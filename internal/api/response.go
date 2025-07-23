package api

import (
	"github.com/labstack/echo/v4"
	"github.com/xjncx/people-info-api/internal/dto"
)

func RespondSuccess[T any](c echo.Context, status int, data T) error {
	return c.JSON(status, dto.SuccessResponse[T]{
		Success: true,
		Data:    data,
	})
}

func RespondError(c echo.Context, status int, message string, details map[string]string) error {
	return c.JSON(status, dto.ErrorResponse{
		Success: false,
		Error:   message,
		Details: details,
	})
}
