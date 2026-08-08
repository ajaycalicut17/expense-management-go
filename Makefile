.PHONY: all run build clean fmt vet tidy tidy-tools migrate migrate-down migrate-reset

TOOLS_MODFILE = -modfile=tools/go.mod

all: fmt vet tidy tidy-tools migrate run

run:
	@echo "Running..."
	@go run ./cmd/server/main.go

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/server ./cmd/server/main.go

clean:
	@echo "Cleaning..."
	@rm -rf bin/

fmt:
	@echo "Formatting..."
	@go fmt ./...

vet:
	@echo "Linting..."
	@go vet ./...

tidy:
	@echo "Tidying Go modules..."
	@go mod tidy

tidy-tools:
	@echo "Tidying tools Go modules..."
	@go mod tidy $(TOOLS_MODFILE)

migrate:
	@echo "Validate, fix and migrate..."
	@go tool $(TOOLS_MODFILE) goose validate
	@go tool $(TOOLS_MODFILE) goose fix
	@go tool $(TOOLS_MODFILE) goose up

migrate-down:
	@echo "Rolling back last migration..."
	@go tool $(TOOLS_MODFILE) goose down

migrate-reset:
	@echo "Resetting all migrations..."
	@go tool $(TOOLS_MODFILE) goose reset