.PHONY: build

CMD_ARGS?=$(filter-out $@, $(MAKECMDGOALS))

up:
	docker compose up -d

down:
	docker compose down --remove-orphans

restart: down up

build:
	docker compose build

workspace-use:
	go work use $(CMD_ARGS)

#-------------------------------------------------------------------------------
# TOOLS:
#-------------------------------------------------------------------------------

tool-docker-config:
	docker compose config

tool-docker-network:
	docker network inspect fine

tool-docker-logs:
	docker compose logs -f $(CMD_ARGS)

tool-docker-prune:
	docker image prune -f
	docker volume prune -f
	docker network prune -f
