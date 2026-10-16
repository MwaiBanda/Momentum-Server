#!/usr/bin/env bash

set -euo pipefail

compose_file="docker-compose.test.yml"
database_url="postgres://momentum_test:momentum_test@127.0.0.1:55432/momentum_test?sslmode=disable"

cleanup() {
  docker compose -f "$compose_file" down --volumes --remove-orphans
}

trap cleanup EXIT

docker compose -f "$compose_file" up --detach --wait

MOMENTUM_TEST_DATABASE_URL="$database_url" \
  go test -mod=mod -run '^TestRepositoryIntegration' \
  ./infrastructure/persistence/repository/user \
  ./infrastructure/persistence/repository/volunteer_service \
  ./infrastructure/persistence/repository/message \
  ./infrastructure/persistence/repository/transaction
