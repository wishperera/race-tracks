.DEFAULT_GOAL := build

lint: ## Perform linting
	golangci-lint run

fmt: ## format project
	go fmt ./...

test: ## Run unit tests
	go test  -mod=vendor ./...  -race

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor  -a -installsuffix cgo -o ./build/race-track-linux-amd64 ./cmd/main.go
    CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 GO111MODULE=on go build -mod=vendor  -a -installsuffix cgo -o ./build/race-track-darwin-amd64 ./cmd/main.go

.PHONY: lint test build fmt