package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/xjncx/people-info-api/internal/client"
	"github.com/xjncx/people-info-api/internal/model"
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

func (s *PersonService) CreatePerson(ctx context.Context, person *model.Person) error {
	var wg sync.WaitGroup

	for _, client := range s.enrichClients {
		wg.Add(1)

		go func() {
			defer wg.Done()
			if err := client.Enrich(ctx, person.FirstName, person); err != nil {
				logger.Log.Info("enrichment failed for %s: %v", client.Name(), err)
			}
		}()
	}

	wg.Wait()

	if err := s.repo.Create(ctx, person); err != nil {
		return fmt.Errorf("create person '%s %s': %w",
			person.FirstName, person.LastName, err)
	}

	return nil
}
