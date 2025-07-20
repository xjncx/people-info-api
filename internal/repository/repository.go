package repository

import (
	"github.com/xjncx/people-info-api/internal/model"
)

type Repository interface {
	Add(person *model.Person) error
}
