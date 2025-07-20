package service

import (
	"context"

	"github.com/xjncx/people-info-api/internal/model"
)

type PersonService struct {
	repo PersonRepository
}

func NewPersonService(r PersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) FindByLastName(ctx context.Context, lastName string) ([]model.Person, error) {
	// пока заглушка
	return nil, nil
}
