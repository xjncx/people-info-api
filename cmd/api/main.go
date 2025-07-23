// @title People Information API
// @version 1.0
// @description API для получения информации о человеке (возраст, пол, национальность)
// @host localhost:8081
// @BasePath /
package main

import (
	"github.com/xjncx/people-info-api/pkg/logger"
	"go.uber.org/zap"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xjncx/people-info-api/internal/api"
	"github.com/xjncx/people-info-api/internal/client"
	"github.com/xjncx/people-info-api/internal/config"
	"github.com/xjncx/people-info-api/internal/repository/pg"
	"github.com/xjncx/people-info-api/internal/service"
)

func main() {

	if err := logger.Init(); err != nil {
		logger.Log.Fatal("Logger initialization failed", zap.Error(err))
	}
	defer logger.Sync()

	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatal("Config loading failed", zap.Error(err))
	}

	logger.Log.Info("Logger initialized")

	db, err := pg.NewDB(cfg)

	if err != nil {
		logger.Log.Fatal("Database connection failed",
			zap.Error(err),
			zap.String("db_host", cfg.DBHost),
			zap.String("db_name", cfg.DBName),
		)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.Log.Error("Database close error", zap.Error(err))
		}
	}()
	logger.Log.Info("Database connected",
		zap.String("host", cfg.DBHost),
		zap.String("name", cfg.DBName),
	)

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
	logger.Log.Info("Starting server",
		zap.String("address", serverAddr),
		zap.String("env", cfg.Environment),
	)

	go func() {
		if err := e.Start(serverAddr); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server shutdown error", zap.Error(err))
	}

	logger.Log.Info("Server exited properly")
}
