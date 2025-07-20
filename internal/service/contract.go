package service

import (
	"context"

	"github.com/xjncx/people-info-api/internal/model"
)

type PersonRepository interface {
	FindByLastName(ctx context.Context, lastName string) ([]model.Person, error)
}
