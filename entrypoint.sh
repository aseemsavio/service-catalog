#!/bin/sh
set -e

# Wait for Postgres to be ready
/app/scripts/wait-for.sh postgres:5432

# Start the API
./service-catalog
