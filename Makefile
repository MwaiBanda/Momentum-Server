.DEFAULT_GOAL := help

GO ?= go
GOFLAGS ?= -mod=mod
SWAG_VERSION ?= v1.16.6

.PHONY: help run build test test-postgres swagger openapi openapi-check cms-api-types cms-api-types-check cms-api-schemas cms-api-schemas-check fmt check docker-build cloudrun-build

help: ## Show available commands
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  %-16s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Start the API locally
	$(GO) run .

build: ## Build the API binary
	$(GO) build -o ./out/dist .

test: ## Run the Go test suite
	$(GO) test $(GOFLAGS) ./...

test-postgres: ## Run PostgreSQL integration tests
	bash scripts/test-postgres.sh

swagger: ## Regenerate Swagger documentation
	$(GO) run github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION) init --parseDependency --parseInternal

openapi: swagger ## Generate the OpenAPI 3 document from Swagger
	$(GO) run ./tools/openapi --input docs/swagger.json --output docs/openapi.json

openapi-check: openapi ## Verify generated API documents are current
	git diff --exit-code -- docs/docs.go docs/swagger.json docs/swagger.yaml docs/openapi.json

cms-api-types: openapi ## Generate CMS transport types from OpenAPI
	yarn --cwd cms api:types

cms-api-types-check: cms-api-types ## Verify CMS transport types are current
	git diff --exit-code -- cms/src/infrastructure/http/generated/openapi.ts

cms-api-schemas: openapi ## Generate CMS runtime schemas from OpenAPI
	yarn --cwd cms api:schemas

cms-api-schemas-check: cms-api-schemas ## Verify CMS runtime schemas are current
	git diff --exit-code -- cms/src/infrastructure/http/generated/schemas.ts

fmt: ## Format Go source files
	gofmt -w $$(find . -type f -name '*.go' -not -path './cms/*' -not -path './vendor/*')

check: ## Check whitespace errors in the Git diff
	git diff --check

docker-build: ## Build the API Docker image
	docker build -t $${IMAGE:-momentum-server} .

cloudrun-build: ## Build the API and Cloud Run sidecar images
	docker build -t $${IMAGE:-momentum-server} .
	docker build -f cloudrun/caddy.Dockerfile -t $${CLOUDRUN_IMAGE_PREFIX:-momentum-server}-caddy .
	docker build -f cloudrun/otel-collector.Dockerfile -t $${CLOUDRUN_IMAGE_PREFIX:-momentum-server}-otel-collector .
	docker build -f cloudrun/prometheus.Dockerfile -t $${CLOUDRUN_IMAGE_PREFIX:-momentum-server}-prometheus .
	docker build -f cloudrun/grafana.Dockerfile -t $${CLOUDRUN_IMAGE_PREFIX:-momentum-server}-grafana .
