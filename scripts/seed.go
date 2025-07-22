package main

import (
	"context"

	"github.com/xjncx/people-info-api/internal/client"
	"github.com/xjncx/people-info-api/internal/config"
	"github.com/xjncx/people-info-api/internal/model"
	"github.com/xjncx/people-info-api/internal/repository/pg"
	"github.com/xjncx/people-info-api/internal/service"
)

var seedData = []struct {
	FirstName  string
	LastName   string
	MiddleName string
}{
	{"Ivan", "Ivanov", "Ivanovich"},
	{"Maria", "Petrova", "Sergeevna"},
	{"Alexander", "Smirnov", "Petrovich"},
	{"Elena", "Kuznetsova", "Andreevna"},
	{"Dmitry", "Popov", "Alexeevich"},
	{"Olga", "Sokolova", "Mikhailovna"},
	{"Sergey", "Novikov", "Vladimirovich"},
	{"Anna", "Morozova", "Nikolaevna"},
	{"Mikhail", "Volkov", "Dmitrievich"},
	{"Natalia", "Fedorova", "Igorevna"},
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.logger.Log.Fatal("Failed to load config: %v", err)
	}

	db, err := pg.NewDB(cfg)

	if err != nil {
		logger.Log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	logger.Log.Info("Connected to database")

	repo := pg.NewPersonRepository(db)

	agifyClient := client.NewAgifyClient(cfg)
	genderizeClient := client.NewGenderizeClient(cfg)
	nationalizeClient := client.NewNationalizeClient(cfg)

	svc := service.NewPersonService(
		repo,
		agifyClient,
		genderizeClient,
		nationalizeClient,
	)

	ctx := context.Background()

	for _, data := range seedData {
		person := &model.Person{
			FirstName:  data.FirstName,
			LastName:   data.LastName,
			MiddleName: data.MiddleName,
		}

		if err := svc.CreatePerson(ctx, person); err != nil {
			logger.Log.Info("Failed to create %s %s: %v", data.FirstName, data.LastName, err)
			continue
		}

		logger.Log.Info("Created: %s %s (age: %d, gender: %s, nationality: %s)",
			person.FirstName, person.LastName, person.Age, person.Gender, person.Nationality)
	}

	logger.Log.Info("Seeding completed!")
}
