package api

import (
	"github.com/xjncx/people-info-api/internal/dto"
	"github.com/xjncx/people-info-api/internal/model"
)

func toPersonResponse(p *model.Person) dto.PersonResponse {
	return dto.PersonResponse{
		ID:          p.ID,
		FirstName:   p.FirstName,
		LastName:    p.LastName,
		MiddleName:  p.MiddleName,
		Age:         p.Age,
		Gender:      p.Gender,
		Nationality: p.Nationality,
	}
}

func toPersonResponses(persons []model.Person) []dto.PersonResponse {
	responses := make([]dto.PersonResponse, len(persons))
	for i, p := range persons {
		responses[i] = toPersonResponse(&p)
	}
	return responses
}

func toPersonModel(req dto.CreatePersonRequest) *model.Person {
	return &model.Person{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
	}
}
