MAKEFILE_DIR := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
WORKSPACE := $(basename $(notdir $(MAKEFILE_DIR)))

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: ## Run linter
	shellcheck $(MAKEFILE_DIR)/src/*

.PHONY: test
test: ## Run the tests
	$(MAKEFILE_DIR)/test/bats/bin/bats $(MAKEFILE_DIR)/test/test.bats

VERSION ?= $(shell git describe --tags --abbrev=0)
VERSION := $(shell echo $(VERSION) | sed -e 's/^v//')
REVISION ?= $(shell git rev-parse --short HEAD)$(shell git diff --quiet || echo -dirty)

.PHONY: build
build: ## Build the debian package
	mkdir -p $(MAKEFILE_DIR)/dist/usr/bin
	cp $(MAKEFILE_DIR)/src/tenminutevpn.bash $(MAKEFILE_DIR)/dist/usr/bin/tenminutevpn
	chmod +x $(MAKEFILE_DIR)/dist/usr/bin/tenminutevpn
	export VERSION=$(VERSION) && \
		export REVISION=$(REVISION) && \
		envsubst < $(MAKEFILE_DIR)/dist/DEBIAN/control.template > $(MAKEFILE_DIR)/dist/DEBIAN/control
	dpkg-deb --build --root-owner-group $(MAKEFILE_DIR)/dist $(MAKEFILE_DIR)/tenminutevpn-$(VERSION)-$(REVISION).deb
	sha256sum *.deb > SHA256SUMS

.PHONY: shell
shell: ## Start the shell (devcontainer)
	docker build -t tenminutevpn:devcontainer -f $(MAKEFILE_DIR)/.devcontainer/Dockerfile $(MAKEFILE_DIR)
	docker run -it --rm -v $(MAKEFILE_DIR):/workspace/$(WORKSPACE) -w /workspace/$(WORKSPACE) tenminutevpn:devcontainer

.PHONY: clean
clean: ## Clean up
	git clean -fdx
