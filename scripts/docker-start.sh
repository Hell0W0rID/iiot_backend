#!/bin/bash

# Docker startup script for IIoT Backend services
# This script determines which service to start based on SERVICE_NAME environment variable

set -e

SERVICE_NAME=${SERVICE_NAME:-"gateway"}
PORT=${PORT:-8000}

echo "Starting IIoT Backend Service: $SERVICE_NAME on port $PORT"

# Wait for dependencies
echo "Waiting for PostgreSQL..."
while ! pg_isready -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" >/dev/null 2>&1; do
    sleep 1
done
echo "PostgreSQL is ready"

echo "Waiting for Redis..."
while ! redis-cli -h "$REDIS_HOST" -p "$REDIS_PORT" ping >/dev/null 2>&1; do
    sleep 1
done
echo "Redis is ready"

# Run migrations only once from metadata service
if [ "$SERVICE_NAME" = "metadata" ]; then
    echo "Running database migrations..."
    go run scripts/migrate.go
    echo "Migrations completed"
fi

# Start the specific service
echo "Starting $SERVICE_NAME service..."
exec ./iiot-backend