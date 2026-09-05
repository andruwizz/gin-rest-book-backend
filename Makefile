GOPATH ?= $(HOME)/go

ifeq ($(OS), Windows_NT)
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
else
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
endif

ifneq (,$(wildcard ./.env))
	include .env
	export
endif

run-api:
	godotenv -f .env go run cmd/main.go

run-test:
	go test ./test

api-docs:
	swag init -g internal/config/swaggo.go --parseDependency

migration $$(enter):
	@read -p "Migration name:" migration_name; \
	migrate create -ext sql -dir database/migrations $$migration_name

migration-up:
	migrate -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_PUBLIC_HOST}:${DATABASE_PUBLIC_PORT})/${DATABASE_NAME}" -path database/migrations up

migration-down:
	migrate -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_PUBLIC_HOST}:${DATABASE_PUBLIC_PORT})/${DATABASE_NAME}" -path database/migrations down