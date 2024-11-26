include .env
export $(shell sed 's/=.*//' .env)

run:
	go run ./cmd/http_server/main.go

migrate:
	migrate -path ./migration -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" up