# TODO build

default: build

HASH := $(shell git rev-parse --short=10 HEAD)

.PHONY: build
build:
	echo "build"

.PHONY: openapi
openapi:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate types -o internal/task/infrastructure/http/api/types.gen.go -package api api/openapi/task.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate chi-server -o internal/task/infrastructure/http/api/server.gen.go -package api api/openapi/task.yaml

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2 run --fix cmd/... internal/...