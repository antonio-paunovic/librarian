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

SERVER_PORT=44455
SLEEP_DURATION=0.5

echo "Building the server:"
go build -o test-server ../cmd/librarian-server

echo "Running the server:"
DB_PATH="../tests/librarian.db"  ./test-server --port=$SERVER_PORT &
sleep $SLEEP_DURATION

echo "Get empty book table:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT books GetBooks
sleep $SLEEP_DURATION

echo "Put a book inside:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT books PostBooks --book.isbn=1234567890123 --book.title=randomBook
sleep $SLEEP_DURATION

echo "Get this book:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT books GetBooks
sleep $SLEEP_DURATION

echo "Update the book:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT books PutBooksIsbn --isbn=1234567890123 --book.isbn=1234567890123 --book.author=randomAuthor
sleep $SLEEP_DURATION

echo "Get updated book:"
go run ../cmd/cli/main.go  --hostname localhost:$SERVER_PORT books GetBooks

echo "Cleanup test-server..."
pkill test-server
rm ./test-server

cleanup
