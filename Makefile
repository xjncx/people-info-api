APP_NAME=people-info-api
DOCKER_COMPOSE=docker-compose

.PHONY: up down restart build seed migrate logs reset

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

build:
	$(DOCKER_COMPOSE) build

restart: down up

logs:
	$(DOCKER_COMPOSE) logs -f app

migrate:
	$(DOCKER_COMPOSE) exec app go run migrate/migrate.go

seed:
	$(DOCKER_COMPOSE) exec app go run scripts/seed.go

reset:
	$(DOCKER_COMPOSE) down -v
	$(DOCKER_COMPOSE) up --build -d

ps:
	$(DOCKER_COMPOSE) ps
