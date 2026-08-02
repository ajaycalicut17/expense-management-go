.PHONY: all run build clean fmt vet tidy

all: fmt vet tidy clean

run:
	@echo "Running..."
	go run ./cmd/server/main.go

build:
	@echo "Building..."
	@mkdir -p bin
	go build -o bin/server ./cmd/server/main.go

clean:
	@echo "Cleaning..."
	rm -rf bin/

fmt:
	@echo "Formatting..."
	go fmt ./...

vet:
	@echo "Linting..."
	go vet ./...

tidy:
	@echo "Tidying Go modules..."
	go mod tidy