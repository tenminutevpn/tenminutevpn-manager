MAKEFILE_DIR := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
WORKSPACE := $(basename $(notdir $(MAKEFILE_DIR)))

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

VERSION ?= $(shell git describe --tags --abbrev=0)
VERSION := $(shell echo $(VERSION) | sed -e 's/^v//')
REVISION ?= $(shell git rev-parse --short HEAD)$(shell git diff --quiet || echo -dirty)

.PHONY: test
test: ## Run the tests
	$(MAKEFILE_DIR)/test/bats/bin/bats $(MAKEFILE_DIR)/test/test.bats

.PHONY: build-%
build-%: ## Build binary
	cd $(MAKEFILE_DIR)/tenminutevpn-manager && \
		export GOOS=linux && \
		export GOARCH=$* && \
		go get -d -v && \
		go build -o $(MAKEFILE_DIR)/tenminutevpn-manager-linux-$* .
	chmod +x $(MAKEFILE_DIR)/tenminutevpn-manager-linux-$*

.PHONY: package-%
package-%: ## Build the Debian package for the given architecture
	mkdir -p $(MAKEFILE_DIR)/dist/usr/bin
	cp $(MAKEFILE_DIR)/tenminutevpn-manager-linux-$* $(MAKEFILE_DIR)/dist/usr/bin/tenminutevpn-manager

	export ARCH=$* && \
		export VERSION=$(VERSION) && \
		export REVISION=$(REVISION) && \
		envsubst < $(MAKEFILE_DIR)/dist/DEBIAN/control.template > $(MAKEFILE_DIR)/dist/DEBIAN/control
	dpkg-deb --build --root-owner-group $(MAKEFILE_DIR)/dist $(MAKEFILE_DIR)/tenminutevpn-manager-$(VERSION)-$(REVISION)-$*.deb

SHA256SUMS:
	sha256sum tenminutevpn-manager-linux-* > SHA256SUMS
	sha256sum tenminutevpn-manager-$(VERSION)-$(REVISION)-*.deb >> SHA256SUMS

.PHONY: checksum
checksum: SHA256SUMS ## Generate the checksums

.PHONY: package
package: package-amd64 package-arm64 ## Build the Debian packages for all architectures

.PHONY: shell
shell: ## Start the shell (devcontainer)
	docker build -t tenminutevpn:devcontainer -f $(MAKEFILE_DIR)/.devcontainer/Dockerfile $(MAKEFILE_DIR)
	docker run -it --rm -v $(MAKEFILE_DIR):/workspace/$(WORKSPACE) -w /workspace/$(WORKSPACE) tenminutevpn:devcontainer

.PHONY: clean
clean: ## Clean up
	git clean -fdx
