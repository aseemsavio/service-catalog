#!/usr/bin/env sh
# wait-for.sh host:port -- command...
set -e

HOSTPORT="$1"
shift
HOST=$(echo "$HOSTPORT" | cut -d: -f1)
PORT=$(echo "$HOSTPORT" | cut -d: -f2)

echo "Waiting for $HOST:$PORT..."
while ! nc -z "$HOST" "$PORT"; do
sleep 0.5
done
echo "Postgres is up."

exec "$@"
