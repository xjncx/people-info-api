package model

import "github.com/google/uuid"

type Person struct {
	ID          uuid.UUID `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	MiddleName  string    `json:"middle_name" db:"middle_name"`
	Age         int       `json:"age" db:"age"`
	Gender      string    `json:"gender" db:"gender"`
	Nationality string    `json:"nationality" db:"nationality"`
}
