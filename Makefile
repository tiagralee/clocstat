all: lint build
build: ## Build clocstat
	go build ./...
lint: ## Run golangci-lint
	golangci-lint run ./...
