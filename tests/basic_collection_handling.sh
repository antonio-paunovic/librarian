#!/bin/bash

set -e

# Initialize database
sh init_test_db.sh

SERVER_PORT=44455
SLEEP_DURATION=0.5

cleanup() {
    echo "Cleanup test-server..."
    pkill test-server
    rm ./test-server
}

trap cleanup EXIT

echo "Building the server:"
go build -o test-server ../cmd/librarian-server

echo "Running the server:"
DB_PATH="../tests/librarian.db"  ./test-server --port=$SERVER_PORT &
sleep $SLEEP_DURATION

echo "Get empty collection table:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT collections GetCollections
sleep $SLEEP_DURATION

echo "Put a collection inside:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT collections PostCollections --collection.name="randomCollection" --body='{"isbn":["0987654321098","9780395082515"]}'
sleep $SLEEP_DURATION

echo "Get this collection:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT collections GetCollections
sleep $SLEEP_DURATION

echo "Delete this collection:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT collections DeleteCollections --name="randomCollection"
sleep $SLEEP_DURATION

cleanup
