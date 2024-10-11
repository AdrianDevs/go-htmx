# Variables
APP_NAME := go-htmx
VERSION := 0.0.1
GO_FILES := $(shell find . -name '*.go')
BUILD_DIR := ./bin
MAIN_FILE := ./main.go

# Build settings
GO_BUILD := go build
GO_RUN := go run
GO_TEST := go test
GO_CLEAN := go clean
TEMPLE_BUILD := templ generate
COMPOSE_UP := docker compose up --detach
COMPOSE_DOWN := docker compose down
AIR_START := air

# Default rule to build the binary
all: build

# Build the application
build: $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(TEMPLE_BUILD)
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "Build complete!"

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	$(COMPOSE_UP)
	$(BUILD_DIR)/$(APP_NAME)

# Run Go tests
test:
	@echo "Running tests..."
	$(GO_TEST) ./...

# Clean the build
clean:
	@echo "Cleaning build..."
	$(GO_CLEAN)
	$(COMPOSE_DOWN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$ go install github.com/a-h/templ/cmd/templ@latest
	$ go install github.com/air-verse/air@latest
	go mod tidy

# Remove Go mod cache
mod-clean:
	@echo "Cleaning module cache..."
	go clean -modcache

# Add a rule for versioning
version:
	@echo "$(APP_NAME) version $(VERSION)"

# Run the application and restart on file changes
watch:
	@echo "Watching $(APP_NAME)..."
	$(COMPOSE_UP)
	$(TEMPLE_BUILD)
	$(AIR_START)

# Help command to display usage of the Makefile
help:
	@echo "Usage:"
	@echo "  make          - Build the application"
	@echo "  make run      - Start containers and run the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean the build and remove containers"
	@echo "  make deps     - Install dependencies"
	@echo "  make version  - Show application version"
	@echo "  make watch    - Start containers, run the application and restart on file changes"
	@echo "  make help     - Show this help message"