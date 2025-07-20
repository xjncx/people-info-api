package api

import (
	"github.com/xjncx/people-info-api/internal/service"
)

type Handler struct {
	PersonService *service.PersonService
}
