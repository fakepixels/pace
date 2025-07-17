# Pace CLI Makefile
.PHONY: help build run clean test install dev deps fmt lint vet release

# Default target
help: ## Show this help message
	@echo "Pace CLI - Development Commands"
	@echo "Usage: make [target]"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Dependencies
deps: ## Download and verify dependencies
	@echo "📦 Installing dependencies..."
	go mod download
	go mod verify
	go mod tidy

# Development
dev: deps ## Run the application in development mode
	@echo "🚀 Starting Pace CLI in development mode..."
	go run main.go

serve: deps ## Run the SSH server for development
	@echo "🌐 Starting Pace CLI SSH server..."
	go run main.go --serve

# Building
build: deps ## Build the binary for current platform
	@echo "🔨 Building Pace CLI..."
	go build -ldflags="-s -w -X main.version=dev" -o bin/pace main.go

build-all: deps ## Build binaries for all platforms
	@echo "🔨 Building for all platforms..."
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=dev" -o bin/pace-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X main.version=dev" -o bin/pace-linux-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X main.version=dev" -o bin/pace-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.version=dev" -o bin/pace-darwin-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X main.version=dev" -o bin/pace-windows-amd64.exe main.go

# Installation
install: build ## Install the binary to $GOPATH/bin or /usr/local/bin
	@echo "📦 Installing Pace CLI..."
	@if [ -n "$(GOPATH)" ] && [ -d "$(GOPATH)/bin" ]; then \
		cp bin/pace $(GOPATH)/bin/; \
		echo "✅ Installed to $(GOPATH)/bin/pace"; \
	elif [ -w "/usr/local/bin" ]; then \
		cp bin/pace /usr/local/bin/; \
		echo "✅ Installed to /usr/local/bin/pace"; \
	else \
		echo "❌ Could not install. Try: sudo make install"; \
		exit 1; \
	fi

install-dev: ## Install from local source
	@echo "📦 Installing from source..."
	go install -ldflags="-s -w -X main.version=dev" .

# Code quality
fmt: ## Format Go code
	@echo "🎨 Formatting code..."
	go fmt ./...

lint: ## Run golangci-lint (requires golangci-lint to be installed)
	@echo "🔍 Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

vet: ## Run go vet
	@echo "🔍 Running go vet..."
	go vet ./...

test: ## Run tests
	@echo "🧪 Running tests..."
	go test -v ./...

# Maintenance
clean: ## Clean build artifacts
	@echo "🧹 Cleaning up..."
	rm -rf bin/
	rm -rf dist/
	go clean

# Release
release: ## Create a release with GoReleaser (requires GITHUB_TOKEN)
	@echo "🚀 Creating release..."
	@if [ -z "$(GITHUB_TOKEN)" ]; then \
		echo "❌ GITHUB_TOKEN is required for release"; \
		exit 1; \
	fi
	goreleaser release --clean

release-snapshot: ## Create a snapshot release (no upload)
	@echo "📸 Creating snapshot release..."
	goreleaser release --snapshot --clean

# Check installation of required tools
check-tools: ## Check if required development tools are installed
	@echo "🔧 Checking development tools..."
	@command -v go >/dev/null 2>&1 || (echo "❌ Go is not installed" && exit 1)
	@echo "✅ Go: $(shell go version)"
	@command -v goreleaser >/dev/null 2>&1 && echo "✅ GoReleaser: $(shell goreleaser --version | head -n1)" || echo "⚠️  GoReleaser not found (optional)"
	@command -v golangci-lint >/dev/null 2>&1 && echo "✅ golangci-lint: $(shell golangci-lint version)" || echo "⚠️  golangci-lint not found (optional)"

# Quick setup for new developers
setup: deps check-tools ## Complete setup for new developers
	@echo "🎉 Development setup complete!"
	@echo "Try running: make dev"