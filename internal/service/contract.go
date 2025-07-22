package service

import (
	"context"

	"github.com/xjncx/people-info-api/internal/model"
)

type PersonRepository interface {
	Create(ctx context.Context, person *model.Person) error
	FindByLastName(ctx context.Context, lastName string) ([]model.Person, error)
}
