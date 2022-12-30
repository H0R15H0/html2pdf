.PHONY: help
DOCKERCOMPOSE := $(shell docker compose version >/dev/null 2>&1; if [ $$? -eq 0 ]; then echo "docker compose"; else echo "docker-compose"; fi)

TMP_CONTAINER_NAME := tmp_container
PDF_BUILDER_SERVICE := pdf_builder_app
POSTGRESQL_SERVICE := postgresql
DB_NAME := pdf_builder_app

# define standard colors
ifneq (,$(findstring xterm,${TERM}))
	BLACK        := $(shell tput -Txterm setaf 0)
	RED          := $(shell tput -Txterm setaf 1)
	GREEN        := $(shell tput -Txterm setaf 2)
	YELLOW       := $(shell tput -Txterm setaf 3)
	LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
	PURPLE       := $(shell tput -Txterm setaf 5)
	BLUE         := $(shell tput -Txterm setaf 6)
	WHITE        := $(shell tput -Txterm setaf 7)
	RESET := $(shell tput -Txterm sgr0)
else
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	LIGHTPURPLE  := ""
	PURPLE       := ""
	BLUE         := ""
	WHITE        := ""
	RESET        := ""
endif

## Server
up: ## docker compose up
	$(DOCKERCOMPOSE) up
up-d: ## docker compose up -d
	$(DOCKERCOMPOSE) up -d
stop: ## docker compose stop
	$(DOCKERCOMPOSE) stop
down: ## docker compose down
	$(DOCKERCOMPOSE) down

## Attach
api-sh: ## docker compose run --rm pdf_builder_app sh
	$(DOCKERCOMPOSE) run --rm $(PDF_BUILDER_SERVICE) sh
postgres-bash: ## docker compose run --rm postgresql bash 
	$(DOCKERCOMPOSE) exec $(POSTGRESQL_SERVICE) bash
postgres-psql: ## docker compose run --rm postgresql psql 
	$(DOCKERCOMPOSE) exec $(POSTGRESQL_SERVICE) bash -c "psql -U $$DB_USER -d $$DB_NAME"

## Database
db-create: ## create database
	$(DOCKERCOMPOSE) exec $(POSTGRESQL_SERVICE) bash -c "psql -U $$DB_USER -c 'create database $(DB_NAME)'"
db-migrate: ## migrate all
	$(DOCKERCOMPOSE) exec $(PDF_BUILDER_SERVICE) sh -c "migrate -database postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable -path database/migrations up"
db-down: ## down all table
	$(DOCKERCOMPOSE) exec $(PDF_BUILDER_SERVICE) sh -c "migrate -database postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable -path database/migrations down"
db-migrate-create: ## create migration file. pass FILE_NAME
	$(DOCKERCOMPOSE) exec $(PDF_BUILDER_SERVICE) sh -c "migrate create -ext sql -dir database/migrations -seq ${FILE_NAME}"

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = "## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-30s${GREEN}%s${RESET}\n", substr($$1, 1, length($$1)-2), $$2} \
		else if (/^## .*$$/) {printf "  ${LIGHTPURPLE}%s${RESET}\n", $$2} \
		}' $(MAKEFILE_LIST)
