package dto

import (
	"github.com/google/uuid"
)

type CreatePersonRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
}

type PersonResponse struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name,omitempty"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	Nationality string    `json:"nationality"`
}

type PersonListResponse struct {
	People []PersonResponse `json:"people"`
}
