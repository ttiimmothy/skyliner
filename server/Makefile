.PHONY: dev test lint migrate seed build clean

# Development
dev:
	$(HOME)/go/bin/air

# Testing
test:
	go test -v ./...

# Linting
lint:
	golangci-lint run

# Database
migrate:
	go run cmd/server/main.go migrate

# Seed data
seed:
	go run cmd/server/main.go seed

# Build
build:
	go build -o bin/server cmd/server/main.go

# Clean
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run server
run:
	go run cmd/server/main.go

# Install air for hot reload
install-air:
	go install github.com/air-verse/air@latest
