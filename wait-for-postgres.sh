#!/bin/sh
# wait-for-postgres.sh

set -e

until PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
>&2 echo "Postgres make migrations"

make migrate_from_app

exec "$@"