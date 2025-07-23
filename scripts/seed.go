package main

import (
	"context"
	"log"

	"github.com/xjncx/people-info-api/internal/client"
	"github.com/xjncx/people-info-api/internal/config"
	"github.com/xjncx/people-info-api/internal/model"
	"github.com/xjncx/people-info-api/internal/repository/pg"
	"github.com/xjncx/people-info-api/internal/service"
	"github.com/xjncx/people-info-api/pkg/logger"
	"go.uber.org/zap"
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
	{"Petr", "Fedorov", "Igorevich"},
}

func main() {
	if err := logger.Init(); err != nil {
		log.Fatalf("Logger init failed: %v", err)
	}
	defer logger.Sync()
	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatal("Config load failed", zap.Error(err))
	}

	db, err := pg.NewDB(cfg)

	if err != nil {
		logger.Log.Fatal("DB connection failed", zap.Error(err))
	}
	defer db.Close()

	ctx := context.Background()
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

	for _, data := range seedData {
		person := &model.Person{
			FirstName:  data.FirstName,
			LastName:   data.LastName,
			MiddleName: data.MiddleName,
		}

		if err := svc.EnrichPerson(ctx, person); err != nil {
			continue
		}

	}

	logger.Log.Info("Seeding completed!")
}
