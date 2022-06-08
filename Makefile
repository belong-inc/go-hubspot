# General
BIN := $(CURDIR)/.bin
PATH := $(abspath $(BIN)):$(PATH)

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

$(BIN):
	mkdir -p $(BIN)

.PHONY: os
os: ## show os name
	@echo "$(UNAME_OS)"

.PHONY: help
help: ## print help
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Go

.PHONY: mod
mod: ## download Go modules
	go mod download

.PHONY: vendor
vendor: ## make go vendor
	go mod vendor

.PHONY: test
test: ## run go test
	go test ./...

# Install golangci-lint
GOLANGCLI_LINT := $(BIN)/golangci-lint
GOLANGCLI_LINT_VERSION := v1.43.0
$(GOLANGCLI_LINT): | $(BIN) ## Install golangci-lint
	@curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | sh -s -- -b $(BIN) $(GOLANGCLI_LINT_VERSION)
	@chmod +x "$(BIN)/golangci-lint"

.PHONY: lint
lint: | $(GOLANGCLI_LINT) ## run golangci-lint with config .golangcli.yml
	$(BIN)/golangci-lint run --verbose --config=.golangci.yml ./...

# Install gofumpt
# This setting is only available for Mac
GOFMPT := $(BIN)/gofumpt
GOFMPT_VERSION := v0.1.0
ifeq "$(UNAME_OS)" "Darwin"
	GOFMPT_BIN=gofumpt_$(GOFMPT_VERSION)_darwin_amd64
endif

$(GOFMPT): | $(BIN) ## Install gofumpt
	@curl -sSfL "https://github.com/mvdan/gofumpt/releases/download/$(GOFMPT_VERSION)/$(GOFMPT_BIN)" \
		-o "$(BIN)/gofumpt"
	@chmod +x "$(BIN)/gofumpt"

.PHONY: fmt
fmt: | $(GOFMPT) ## format files via gofumpt and list impacted files
	@$(BIN)/gofumpt -l -w . ./legacy
