.PHONY: build

-include build/.env

COMPOSE_FILE:=./build/docker-compose.yml

up:
	docker compose -f $(COMPOSE_FILE) up -d

down:
	docker compose -f $(COMPOSE_FILE) down

restart: down up

build:
	docker compose -f $(COMPOSE_FILE) build

logs:
	docker compose -f $(COMPOSE_FILE) logs -f

# go mod edit -replace=[current package]=[target package]
go-mod-replace:
	go mod edit -replace=$(p)
