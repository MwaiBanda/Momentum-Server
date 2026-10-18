.DEFAULT_GOAL := help

GO ?= go
GOFLAGS ?= -mod=mod
SWAG_VERSION ?= v1.16.6

.PHONY: help run build test test-postgres swagger openapi openapi-check cms-api-types cms-api-types-check fmt check docker-build

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

fmt: ## Format Go source files
	gofmt -w $$(find . -type f -name '*.go' -not -path './cms/*' -not -path './vendor/*')

check: ## Check whitespace errors in the Git diff
	git diff --check

docker-build: ## Build the API Docker image
	docker build -t momentum-server .
