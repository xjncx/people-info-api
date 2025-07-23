package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/xjncx/people-info-api/internal/client"
	"github.com/xjncx/people-info-api/internal/model"
	"github.com/xjncx/people-info-api/pkg/logger"
	"go.uber.org/zap"
)

type PersonService struct {
	repo          PersonRepository
	enrichClients []client.EnrichmentClient
}

func NewPersonService(
	repo PersonRepository,
	clients ...client.EnrichmentClient,
) *PersonService {
	return &PersonService{
		repo:          repo,
		enrichClients: clients,
	}
}

func (s *PersonService) FindByLastName(ctx context.Context, lastName string) ([]model.Person, error) {
	people, err := s.repo.FindByLastName(ctx, lastName)
	if err != nil {
		return nil, fmt.Errorf("find by last name '%s': %w", lastName, err)
	}
	return people, nil
}

func (s *PersonService) EnrichPerson(ctx context.Context, person *model.Person) error {
	//на проде такое решение не подойдет, была идея сделать воркер, который бы раз в секунду
	//брал задачи на обогащение

	var wg sync.WaitGroup
	results := make(chan *client.EnrichmentData, len(s.enrichClients))

	for _, enrichClient := range s.enrichClients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if data, err := enrichClient.Enrich(ctx, person.FirstName); err == nil {
				results <- data
			} else {
				logger.Log.Warn("Enrichment failed",
					zap.String("enricher", enrichClient.Name()),
					zap.Error(err))
			}
		}()
	}

	wg.Wait()
	close(results)

	for data := range results {
		if data.Age != nil {
			person.Age = *data.Age
		}
		if data.Gender != nil {
			person.Gender = *data.Gender
		}
		if data.Nationality != nil {
			person.Nationality = *data.Nationality
		}
	}

	return s.repo.Create(ctx, person)
}
