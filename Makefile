.PHONY: build test clean docker docker-compose-up docker-compose-down lint fmt

# Variables
BINARY_NAME=cardano-monitor
DOCKER_IMAGE=cardano-node-monitor
VERSION=$(shell git describe --tags --always --dirty)

# Build the application
build:
	CGO_ENABLED=0 go build -ldflags="-X main.Version=$(VERSION)" -o $(BINARY_NAME) ./cmd/

# Build for multiple platforms
build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X main.Version=$(VERSION)" -o $(BINARY_NAME)-linux-amd64 ./cmd/
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.Version=$(VERSION)" -o $(BINARY_NAME)-darwin-amd64 ./cmd/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-X main.Version=$(VERSION)" -o $(BINARY_NAME)-windows-amd64.exe ./cmd/

# Run tests
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run benchmarks
bench:
	go test -bench=. -benchmem ./...

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)*
	rm -f coverage.txt

# Build Docker image
docker:
	docker build -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

# Start services with docker-compose
docker-compose-up:
	docker-compose up -d

# Stop services
docker-compose-down:
	docker-compose down

# Lint the code
lint:
	golangci-lint run

# Format the code
fmt:
	go fmt ./...
	goimports -w .

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run the application
run:
	go run ./cmd/ --node-url http://localhost:12798

# Install development tools
install-tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Generate mocks (if using mockgen)
generate:
	go generate ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the application"
	@echo "  build-all      - Build for multiple platforms"
	@echo "  test           - Run tests"
	@echo "  bench          - Run benchmarks"
	@echo "  clean          - Clean build artifacts"
	@echo "  docker         - Build Docker image"
	@echo "  docker-compose-up   - Start services with docker-compose"
	@echo "  docker-compose-down - Stop services"
	@echo "  lint           - Lint the code"
	@echo "  fmt            - Format the code"
	@echo "  deps           - Install dependencies"
	@echo "  run            - Run the application"
	@echo "  install-tools  - Install development tools"
	@echo "  generate       - Generate code"
	@echo "  help           - Show this help"