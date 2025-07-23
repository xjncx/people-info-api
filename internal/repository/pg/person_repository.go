package pg

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx" //выбрал sqlx для тестового, про sqrl знаю
	"github.com/xjncx/people-info-api/internal/model"
)

type PersonRepository struct {
	db *sqlx.DB
}

func NewPersonRepository(db *sqlx.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(ctx context.Context, person *model.Person) error {

	query := `
        INSERT INTO people (first_name, last_name, middle_name, age, gender, nationality)
        VALUES (:first_name, :last_name, :middle_name, :age, :gender, :nationality)
        RETURNING id
    `

	rows, err := r.db.NamedQueryContext(ctx, query, person)
	if err != nil {
		return fmt.Errorf("insert person: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&person.ID)
		if err != nil {
			return fmt.Errorf("scan person id: %w", err)
		}
	}

	return nil

}

func (r *PersonRepository) FindByLastName(ctx context.Context, lastName string) ([]model.Person, error) {

	query := `
        SELECT id, first_name, last_name, age, gender, nationality 
        FROM people 
        WHERE last_name = $1
    `
	rows, err := r.db.QueryContext(ctx, query, lastName)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var people []model.Person
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age, &p.Gender, &p.Nationality); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		people = append(people, p)
	}

	return people, nil
}
