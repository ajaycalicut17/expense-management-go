GOOSE = go run github.com/pressly/goose/v3/cmd/goose@v3.27.3

.PHONY: all run build clean fmt vet tidy

all: fmt vet tidy run

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

migrate:
	@echo "Validate, fix and migrate..."
	@$(GOOSE) validate
	@$(GOOSE) fix
	@$(GOOSE) up

migrate-down:
	@echo "Rolling back last migration..."
	@$(GOOSE) down

migrate-reset:
	@echo "Resetting all migrations..."
	@$(GOOSE) reset