package dto

import (
	"github.com/google/uuid"
)

// Request структуры
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
	//    Emails      []EmailResponse `json:"emails,omitempty"`
}

type PersonListResponse struct {
	People []PersonResponse `json:"people"`
}
