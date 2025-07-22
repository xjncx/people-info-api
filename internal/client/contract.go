package client

import (
	"context"

	"github.com/xjncx/people-info-api/internal/model"
)

type EnrichmentClient interface {
	Enrich(ctx context.Context, name string, person *model.Person) error
	Name() string
}
