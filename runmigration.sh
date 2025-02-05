#!/bin/bash

# Set database connection URL
DB_HOST="localhost:5432/gorssdb"
DB_URL="postgres://root:123@localhost:5432/gorssdb?sslmode=disable"

# Path to migrations folder (change if necessary)
MIGRATIONS_DIR="./sql/schema"

# Ensure goose is installed
if ! command -v goose &> /dev/null
then
    echo "Goose is not installed. Installing..."
    go install github.com/pressly/goose/v3/cmd/goose@latest
fi

# Apply migrations
echo "Applying migrations to database: $DB_HOST"
goose -dir "$MIGRATIONS_DIR" postgres "$DB_URL" up

echo "Migration completed!"