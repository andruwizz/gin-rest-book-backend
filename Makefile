GOPATH ?= $(HOME)/go

ifeq ($(OS), Windows_NT)
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
else
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
endif

run-backend:
	go run cmd/main.go