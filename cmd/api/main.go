package main

import (
	"github.com/xjncx/people-info-api/pkg/logger"

	"github.com/xjncx/people-info-api/internal/api"
	"github.com/xjncx/people-info-api/internal/client"
	"github.com/xjncx/people-info-api/internal/config"
	"github.com/xjncx/people-info-api/internal/repository/pg"
	"github.com/xjncx/people-info-api/internal/service"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		logger.logger.Log.Fatal("Failed to load config: %v", err)
	}

	if err := logger.Init(); err != nil {
		logger.logger.Log.Fatal("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Log.Info("Logger initialized")

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

	personService := service.NewPersonService(
		repo,
		agifyClient,
		genderizeClient,
		nationalizeClient,
	)

	handler := &api.Handler{PersonService: personService}
	e := api.NewRouter(handler)

	serverAddr := ":" + cfg.ServerPort
	logger.Log.Info("Starting server on %s", serverAddr)
	if err := e.Start(serverAddr); err != nil {
		logger.Log.Fatal(err)
	}
}
