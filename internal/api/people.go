package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xjncx/people-info-api/internal/dto"
)

func (h *Handler) CreatePerson(c echo.Context) error {
	var req dto.CreatePersonRequest

	if err := c.Bind(&req); err != nil {
		logger.Log.Info("Decode error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if req.FirstName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "first_name is required")
	}

	person := toPersonModel(req)

	if err := h.PersonService.CreatePerson(c.Request().Context(), person); err != nil {
		logger.Log.Info("Failed to create person: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create person")
	}

	response := toPersonResponse(person)

	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) FindByLastName(c echo.Context) error {
	lastName := c.QueryParam("last_name")

	if lastName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "last_name parameter is required")
	}

	persons, err := h.PersonService.FindByLastName(c.Request().Context(), lastName)
	if err != nil {
		logger.Log.Info("Failed to find persons: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch persons")
	}

	response := dto.PersonListResponse{
		People: toPersonResponses(persons),
	}

	return c.JSON(http.StatusOK, response)
}
