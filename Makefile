default: build-api

HASH := $(shell git rev-parse --short=10 HEAD)

.PHONY: build-api
build-api:
	docker build -t task-tracker/api:$(HASH) -f ./build/package/Dockerfile.api --target release .
	docker tag task-tracker/api:$(HASH) task-tracker/api:latest

.PHONY: up
up:
	@HASH=$(HASH) docker compose -p task-tracker -f build/package/docker-compose-dev.yml up -d

.PHONY: stop
stop:
	@HASH=$(HASH) docker compose -p task-tracker -f build/package/docker-compose-dev.yml stop

.PHONY: down
down:
	@HASH=$(HASH) docker compose -p task-tracker -f build/package/docker-compose-dev.yml down

.PHONY: gen-openapi
gen-openapi:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate types -o internal/task_tracker/infrastructure/http/api/types.gen.go -package api api/openapi/task_tracker.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate chi-server -o internal/task_tracker/infrastructure/http/api/server.gen.go -package api api/openapi/task_tracker.yaml

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2 run --fix cmd/... internal/...
