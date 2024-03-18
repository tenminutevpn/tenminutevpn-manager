MAKEFILE_DIR := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
WORKSPACE := $(basename $(notdir $(MAKEFILE_DIR)))

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

VERSION ?= $(shell git describe --tags --abbrev=0)
VERSION := $(shell echo $(VERSION) | sed -e 's/^v//')
REVISION ?= $(shell git rev-parse --short HEAD)$(shell git diff --quiet || echo -dirty)

.PHONY: build-%
build-%: ## Build the Debian package for the given architecture
	mkdir -p $(MAKEFILE_DIR)/dist/usr/bin

	cd $(MAKEFILE_DIR)/tenminutevpn-manager && \
		GOOS=linux \
		GOARCH=$* \
	go build -o $(MAKEFILE_DIR)/dist/usr/bin/tenminutevpn-manager .
	chmod +x $(MAKEFILE_DIR)/dist/usr/bin/tenminutevpn-manager

	export ARCH=$* && \
		export VERSION=$(VERSION) && \
		export REVISION=$(REVISION) && \
		envsubst < $(MAKEFILE_DIR)/dist/DEBIAN/control.template > $(MAKEFILE_DIR)/dist/DEBIAN/control
	dpkg-deb --build --root-owner-group $(MAKEFILE_DIR)/dist $(MAKEFILE_DIR)/tenminutevpn-$(VERSION)-$(REVISION)-$*.deb
	sha256sum *.deb > SHA256SUMS

.PHONY: build
build: build-amd64 build-arm64 ## Build the Debian packages for all architectures

.PHONY: shell
shell: ## Start the shell (devcontainer)
	docker build -t tenminutevpn:devcontainer -f $(MAKEFILE_DIR)/.devcontainer/Dockerfile $(MAKEFILE_DIR)
	docker run -it --rm -v $(MAKEFILE_DIR):/workspace/$(WORKSPACE) -w /workspace/$(WORKSPACE) tenminutevpn:devcontainer

.PHONY: clean
clean: ## Clean up
	git clean -fdx
