# Use default shell BASH.
SHELL_PATH := /bin/bash
SHELL := /usr/bin/env bash

# ==============================================================================
# Define dependencies

NAME            := go-hubspot
GOBIN           := $$HOME/go/bin
GOVULNCHECK     := $(GOBIN)/govulncheck
STATICCHECK     := $(GOBIN)/staticcheck
TOOLS           := tools

# ==============================================================================
# Defining all make targets

.DEFAULT_GOAL := all

.PHONY: all
all: update fmt lint vulncheck tidy test

.PHONY: fmt
fmt:
	@echo "-- Formatting Go files --"
	gofmt -w -s . && \
	gofmt -w -s legacy && \
	goimports -local github.com/belong-inc/go-hubspot -w .

.PHONY: generate
generate: ## generate go code (e.g. make generate OBJECT=Contact FILEPATH=contact.csv)
	@cd $(TOOLS)/model_generator && go run model_gen.go $(OBJECT) $(FILEPATH)
	$(MAKE) fmt

.PHONY: lint
lint:
	@echo "-- Lint check for Go files --"
	go mod tidy
	golangci-lint run
	CGO_ENABLED=0 go vet ./...

.PHONY: staticcheck
staticcheck:
	$(STATICCHECK) -checks=all ./...

.PHONY: vulncheck
vulncheck:
	$(GOVULNCHECK) ./...

.PHONY: update
update:
	go get -u ./...

.PHONY: tidy
tidy:
	@echo "-- Tidying Go modules --"
	go mod tidy

.PHONY: clean
clean:
	@echo " -- Cleaning Go files --"
	@go clean -cache -testcache -modcache
	@rm $(GO_DIR)/*.zip

.PHONY: build
build:
	@echo "-- Building binaries --"
	go build -v ./...

.PHONY:
test: build
	@echo "-- Running tests --"
	CGO_ENABLED=1 go test ./...
