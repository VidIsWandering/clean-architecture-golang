.PHONY: build test run clean fmt lint

# Build the application
build:
	cd src && go build -o ../bin/task-manager main.go

# Run tests
test:
	cd src && go test ./...

test-unit:
	cd src && go test ./...

test-integration:
	cd src && go test -tags=integration ./...

test-all: test-unit test-integration
	@echo "All tests executed"

# Run the application
run:
	cd src && go run main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Format code
fmt:
	cd src && go fmt ./...

# Lint code
lint:
	cd src && golint ./...

# Run all checks
check: fmt lint test build