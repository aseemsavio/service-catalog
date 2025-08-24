APP := service-catalog

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP) ./cmd/api

integration-test:
	go test ./integration -v

tidy:
	go mod tidy

up:
	docker compose up -d

down:
	docker compose down -v

psql:
	docker-compose exec postgres psql -U $$PG_USER -d $$PG_DB
