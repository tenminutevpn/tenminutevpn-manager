MAKEFILE_DIR := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: ## Run linter
	shellcheck $(MAKEFILE_DIR)/src/*

.PHONY: build
build: ## Build the docker image
	docker build -t tenminutevpn:latest $(MAKEFILE_DIR)

.PHONY: run
run: ## Run the docker image
	docker run -it --rm tenminutevpn:latest

.PHONY: test
test: ## Run the tests
	$(MAKEFILE_DIR)/test/bats/bin/bats $(MAKEFILE_DIR)/test/test.bats
