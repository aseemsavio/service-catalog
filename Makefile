APP := service-catalog

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP) ./cmd/api

test:
	go test ./...

lint:
	@echo "Add golangci-lint if desired"

up:
	docker-compose up -d --build

down:
	docker-compose down -v

logs:
	docker-compose logs -f api

psql:
	docker-compose exec postgres psql -U $$PG_USER -d $$PG_DB

migrate-up:
	@echo "Apply migrations with psql"
	docker-compose exec -T postgres bash -lc 'for f in /migrations/*.sql; do psql -U $$POSTGRES_USER -d $$POSTGRES_DB -f $$f; done'
