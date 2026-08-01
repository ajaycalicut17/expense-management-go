.PHONY: all run fmt vet tidy

all: fmt vet tidy

run:
	@echo "Running..."
	go run ./cmd/server/main.go

fmt:
	@echo "Formatting..."
	go fmt ./...

vet:
	@echo "Linting..."
	go vet ./...

tidy:
	@echo "Tidying Go modules..."
	go mod tidy