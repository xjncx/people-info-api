package client

import (
	"context"
)

type EnrichmentData struct {
	Age         *int
	Gender      *string
	Nationality *string
}

type EnrichmentClient interface {
	Enrich(ctx context.Context, name string) (*EnrichmentData, error)
	Name() string
}
