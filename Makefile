APP_NAME := myserver
VERSION ?= latest

.PHONY: build
build:
	@go build \
		-ldflags "-X main.version=$(VERSION)" \
		-o bin/$(APP_NAME) \
		./cmd/$(APP_NAME)/

.PHONY: fmt
fmt:
	@find . -iname "*.go" -not -path "./vendor/**" | xargs gofmt -s -w

.PHONY: test
test: ## run all test on your local machine.
	go test ./...