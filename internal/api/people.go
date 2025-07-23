package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xjncx/people-info-api/internal/dto"
	"github.com/xjncx/people-info-api/pkg/logger"
	"go.uber.org/zap"
)

// @Summary Поиск людей по фамилии
// @Tags People
// @Accept  json
// @Produce  json
// @Param last_name query string true "Фамилия"
// @Success 200 {object} dto.SuccessResponse{data=dto.PersonListResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /people/search [get]
func (h *Handler) FindByLastName(c echo.Context) error {
	lastName := c.QueryParam("last_name")

	if lastName == "" {
		logger.Log.Warn("Validation failed: empty last_name")
		return RespondError(c, http.StatusBadRequest, "last_name parameter is required", nil)
	}

	persons, err := h.PersonService.FindByLastName(c.Request().Context(), lastName)
	if err != nil {
		logger.Log.Error("FindByLastName failed", zap.Error(err))
		return RespondError(c, http.StatusInternalServerError, "failed to fetch persons", nil)
	}

	response := dto.PersonListResponse{
		People: toPersonResponses(persons),
	}

	return RespondSuccess(c, http.StatusOK, response)
}
