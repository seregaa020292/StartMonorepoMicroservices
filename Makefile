.PHONY: build

# https://stackoverflow.com/a/6273809/1334666
CMD_ARGS?=$(filter-out $@, $(MAKECMDGOALS)) $(MAKEFLAGS)
%:
	@true

up:
	docker compose up -d

down:
	docker compose down --remove-orphans

restart: down up

build:
	docker compose build --pull $(CMD_ARGS)

rebuild:
	docker compose up -d --no-deps --build $(CMD_ARGS)

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
