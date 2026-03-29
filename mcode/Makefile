.PHONY: help install run test build clean lint fmt vet

help:
	@echo "Available commands:"
	@echo "  make install   - Download dependencies"
	@echo "  make run       - Run the server"
	@echo "  make test      - Run tests"
	@echo "  make build     - Build binary"
	@echo "  make clean     - Clean build files"
	@echo "  make lint      - Run linter"
	@echo "  make fmt       - Format code"
	@echo "  make vet       - Check code"

install:
	go mod download
	go mod tidy

run:
	go run cmd/server/main.go

test:
	go test ./cmd/server -v

build:
	go build -o bin/server cmd/server/main.go

clean:
	rm -rf bin/
	go clean

lint:
	go fmt ./...
	go vet ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

# Запуск сервера с автоперезагрузкой (требует air)
watch:
	air
