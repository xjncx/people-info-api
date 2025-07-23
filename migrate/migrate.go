package main

import (
	"log"

	"github.com/pressly/goose/v3"

	"github.com/xjncx/people-info-api/internal/config"
	"github.com/xjncx/people-info-api/internal/repository/pg"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := pg.NewDB(cfg)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("goose dialect: %v", err)
	}

	if err := goose.Up(db.DB, "migrate"); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}
