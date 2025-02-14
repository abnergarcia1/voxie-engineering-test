SHELL := /bin/bash
PREFIX?=$(shell pwd)
IMPORT_CONTAINER=voxie/engineering-test:incoming
GO111MODULE := on

.PHONY: all
all: help

.PHONY: db
db: ## Start local development mysql database on :3306
	@echo "+ $@"
	@docker-compose up -d db

.PHONY: import
import: ## Send an import JSON request to :8000/import
	@echo "+ $@"
	@docker run --rm --network=host ${IMPORT_CONTAINER}

.PHONY: import-linux
import-linux: ## Send an import JSON request to :8000/import on a linux-platform
	@echo "+ $@"
	@docker run --rm --env INCOMING_LINUX=1 --network=host ${IMPORT_CONTAINER}

.PHONY: build
build: ## build the current go project
	cd ./project/cmd/project && go build -v main.go && cd ../..

.PHONY: run
run: ## run the current go project
	cd ./project/cmd/project && go run main.go && cd ../..

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

