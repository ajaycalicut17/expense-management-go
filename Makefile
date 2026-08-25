.PHONY: all run build clean fmt vet tidy tidy-tools migrate-validate migrate-fix migrate-up migrate migrate-down migrate-reset air

TOOLS_MODFILE = -modfile=tools/go.mod

all: fmt vet tidy tidy-tools migrate

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

migrate-validate:
	@echo "Validating migrations..."
	@go tool $(TOOLS_MODFILE) goose validate

migrate-fix:
	@echo "Fixing migrations..."
	@go tool $(TOOLS_MODFILE) goose fix

migrate-up:
	@echo "Running migrations..."
	@go tool $(TOOLS_MODFILE) goose up

migrate: migrate-validate migrate-fix migrate-up

migrate-down:
	@echo "Rolling back last migration..."
	@go tool $(TOOLS_MODFILE) goose down

migrate-reset:
	@echo "Resetting all migrations..."
	@go tool $(TOOLS_MODFILE) goose reset

air:
	@echo "Air Running..."
	@go tool $(TOOLS_MODFILE) air -c tools/.air.toml
