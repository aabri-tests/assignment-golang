default: clean run-server lint test build e2e run-e2e-tests k6

BINARY_NAME=server
SERVER_URL=http://localhost:8080


.PHONY: help
help: ## Print this help with list of available commands/targets and their purpose
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: run-e2e-tests
run-e2e-tests: ## run e2e tests
	@make run-server
	@make e2e
	@make clean

.PHONY: clean
clean:  ## cleanup test cover output
	@echo "Cleaning up..."
	@docker-compose down
	@rm -f $(BINARY_NAME)

.PHONY: run-server
run-server: ## Run the Go REST API and its dependencies in Docker containers
	@echo "Starting the API server ..."
	@docker-compose up -d

.PHONY: e2e
e2e: ## run e2e tests
	@echo "Running e2e tests..."
	cd e2e && \
	go test -v -ginkgo.v -timeout 3600s --coverpkg=github.com/aabri-assignments/assignment-golang/handlers/...

.PHONY: test
test:  ## run unit tests
	@echo "Running tests..."
	@go test -v -race -count=1 $(shell go list ./... | grep -v e2e)

.PHONY: build
build:  ## build the server
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(BINARY_NAME)

.PHONY: lint
lint:  ## run linters on the code base
	golangci-lint run ./...

.PHONY: k6
k6:  ## Simulate multiple requests in the same time using k6
	@make run-server
	k6 run k6.js
	@make clean