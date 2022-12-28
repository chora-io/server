#!/usr/bin/make -f

###############################################################################
###                                Generate                                 ###
###############################################################################

gen:
	@find . -name 'go.mod' -type f -execdir go generate ./... \;

.PHONY: gen

###############################################################################
###                                Database                                 ###
###############################################################################

db:
	@docker-compose down -v --remove-orphans
	@docker-compose up db

.PHONY: db

###############################################################################
###                                   App                                   ###
###############################################################################

app:
	@go run ./cmd/app

.PHONY: app

###############################################################################
###                                  Local                                  ###
###############################################################################

local:
	@docker-compose down -v --remove-orphans
	@docker-compose up --abort-on-container-exit --exit-code-from app

.PHONY: local

###############################################################################
###                              Documentation                              ###
###############################################################################

docs:
	@echo "Wait a few seconds and then visit http://localhost:6060/pkg/github.com/choraio/server/"
	godoc -http=:6060

.PHONY: docs

###############################################################################
###                               Go Modules                                ###
###############################################################################

verify:
	@echo "Verifying all go module dependencies..."
	@find . -name 'go.mod' -type f -execdir go mod verify \;

tidy:
	@echo "Cleaning up all go module dependencies..."
	@find . -name 'go.mod' -type f -execdir go mod tidy \;

.PHONY: verify tidy

###############################################################################
###                             Lint / Format                               ###
###############################################################################

lint:
	@echo "Linting all go modules..."
	@find . -name 'go.mod' -type f -execdir golangci-lint run --out-format=tab \;

lint-fix: format
	@echo "Attempting to fix lint errors in all go modules..."
	@find . -name 'go.mod' -type f -execdir golangci-lint run --fix --out-format=tab --issues-exit-code=0 \;

format_filter = -name '*.go' -type f

format_local = github.com/choraio/server

format:
	@echo "Formatting all go modules..."
	@find . $(format_filter) | xargs gofmt -s -w
	@find . $(format_filter) | xargs goimports -w -local $(subst $(whitespace),$(comma),$(format_local))
	@find . $(format_filter) | xargs misspell -w

.PHONY: lint lint-fix format

###############################################################################
###                                  Tests                                  ###
###############################################################################

CURRENT_DIR=$(shell pwd)
GO_MODULES=$(shell find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)

test: test-all

test-all:
	@for module in $(GO_MODULES); do \
		echo "Testing Module $$module"; \
		cd ${CURRENT_DIR}/$$module; \
		go test ./...; \
	done

test-app:
	@echo "Testing Module ."
	@go test ./... \
		-coverprofile=coverage-app.out -covermode=atomic

test-coverage:
	@cat coverage*.out | grep -v "mode: atomic" >> coverage.txt

test-clean:
	@go clean -testcache
	@find . -name 'coverage.txt' -delete
	@find . -name 'coverage*.out' -delete

.PHONY: test test-all test-app test-coverage test-clean

###############################################################################
###                               Go Version                                ###
###############################################################################

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MIN_GO_MAJOR_VERSION = 1
MIN_GO_MINOR_VERSION = 19
GO_VERSION_ERROR = Golang version $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION) is not supported, \
please update to at least $(MIN_GO_MAJOR_VERSION).$(MIN_GO_MINOR_VERSION)

go-version:
	@echo "Verifying go version..."
	@if [ $(GO_MAJOR_VERSION) -gt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 0; \
	elif [ $(GO_MAJOR_VERSION) -lt $(MIN_GO_MAJOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MIN_GO_MINOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	fi

.PHONY: go-version

###############################################################################
###                                  Tools                                  ###
###############################################################################

tools: go-version
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/client9/misspell/cmd/misspell@latest
	@go install golang.org/x/tools/cmd/goimports@latest

.PHONY: tools

###############################################################################
###                                 Clean                                   ###
###############################################################################

clean: test-clean
	@docker-compose down -v --remove-orphans --rmi all

.PHONY: clean
