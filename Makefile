GOPATH ?= $(HOME)/go

ifeq ($(OS), Windows_NT)
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
else
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
endif

run:
	go run cmd/main.go

migration $$(enter):
	@read -p "Migration name:" migration_name; \
	migrate create -ext sql -dir database/migrations $$migration_name

migration-up:
	migrate -database "mysql://user:password@tcp(127.0.0.1:3306)/app" -path database/migrations up

migration-down:
	migrate -database "mysql://user:password@tcp(127.0.0.1:3306)/app" -path database/migrations down