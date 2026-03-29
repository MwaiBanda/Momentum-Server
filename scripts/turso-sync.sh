#!/usr/bin/env bash
set -euo pipefail

command_name="${1:-}"

read_env_var() {
  local key="$1"
  local current="${!key:-}"

  if [[ -n "$current" ]]; then
    printf '%s' "$current"
    return
  fi

  if [[ -f .env ]]; then
    local from_file
    from_file="$(grep "^${key}=" .env | head -n 1 | cut -d= -f2- || true)"
    printf '%s' "$from_file"
    return
  fi

  printf ''
}

turso_database_url="$(read_env_var TURSO_DATABASE_URL)"
turso_auth_token="$(read_env_var TURSO_AUTH_TOKEN)"

if [[ -z "$turso_database_url" || -z "$turso_auth_token" ]]; then
  echo "Missing TURSO_DATABASE_URL or TURSO_AUTH_TOKEN (export them or add to .env)."
  exit 1
fi

connection_url="${turso_database_url}?authToken=${turso_auth_token}"

case "$command_name" in
  push)
    if ! awk '!/^BEGIN TRANSACTION;$/ && !/^COMMIT;$/' prisma/turso-sync.sql | turso db shell "$connection_url"; then
      echo "Push failed. If the remote schema already exists, run: npm run db:sync-turso:reset"
      exit 1
    fi
    ;;
  reset)
    {
      echo "PRAGMA foreign_keys=OFF;"
      echo "DROP TABLE IF EXISTS \"Message\";"
      echo "DROP TABLE IF EXISTS \"Passage\";"
      echo "DROP TABLE IF EXISTS \"Service\";"
      echo "DROP TABLE IF EXISTS \"Note\";"
      echo "DROP TABLE IF EXISTS \"User\";"
      echo "DROP TABLE IF EXISTS \"VolunteerService\";"
      echo "DROP TABLE IF EXISTS \"VolunteeredServiceDay\";"
      echo "DROP TABLE IF EXISTS \"VolunteeredServiceParticipant\";"
      echo "DROP TABLE IF EXISTS \"Meal\";"
      echo "DROP TABLE IF EXISTS \"VolunteeredMeal\";"
      echo "DROP TABLE IF EXISTS \"MealUpdates\";"
      echo "DROP TABLE IF EXISTS \"MealParticipant\";"
      echo "DROP TABLE IF EXISTS \"Transaction\";"
    } | turso db shell "$connection_url"
    ;;
  *)
    echo "Usage: bash scripts/turso-sync.sh [push|reset]"
    exit 1
    ;;
esac
