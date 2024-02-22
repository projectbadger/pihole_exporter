# Golang Makefile for package golang-stress-tester

VERSION ?= $(shell grep "const Version" config/config.go | awk -F'= ' '{print $2}' | sed 's/"//g')
PACKAGE_NAME ?= $(shell grep "^module " go.mod | sed 's/^module\s//')
RUN_ARGS ?= -h
BUILD_MAIN_FILE = .
GO_CMD = go
GOPATH ?= $(shell $(GO_CMD) env GOPATH)
GOAUTODOC_CMD = $(GOPATH)/bin/autodoc
ECHO = "/usr/bin/echo"
# GOAUTODOC_CMD = $(shell go env GOPATH)/bin/goautodoc
TIME = $(shell date)
# Linker-added variables (residing in the config/ module)
LDFLAGS = -X '$(PACKAGE_NAME)/config.PackageName=$(PACKAGE_NAME)' -X '$(PACKAGE_NAME)/config.Version=$(VERSION)' -X '$(PACKAGE_NAME)/config.BuildTime=$(TIME)'

RUN_PORT ?= 9311

test:
	go test ./...

get-pkg:
	# go install github.com/projectbadger/autodoc@latest
	go mod tidy

setup-install:
	$(GOAUTODOC_CMD) -version || go install github.com/projectbadger/autodoc@latest
	$(GOAUTODOC_CMD) mod tidy

setup: setup-install

build-linux-amd64:
	echo "Compiling linux-amd64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/amd64/$(PACKAGE_NAME)-linux-amd64 $(BUILD_MAIN_FILE)

build-linux-arm64:
	echo "Compiling linux-arm64"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/arm64/$(PACKAGE_NAME)-linux-arm64 $(BUILD_MAIN_FILE)

build-linux-arm:
	echo "Compiling linux-arm"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/arm/$(PACKAGE_NAME)-linux-arm $(BUILD_MAIN_FILE)

build-windows-amd64:
	echo "Compiling linux-arm64"
	CGO_ENABLED=0 GOOS=windows $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/windows/amd64/$(PACKAGE_NAME)-windows-amd64 $(BUILD_MAIN_FILE)

build-darwin-amd64:
	echo "Compiling linux-amd64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/darwin/amd64/$(PACKAGE_NAME)-darwin-amd64 $(BUILD_MAIN_FILE)

build-darwin-arm64:
	echo "Compiling linux-amd64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/darwin/arm64/$(PACKAGE_NAME)-darwin-arm64 $(BUILD_MAIN_FILE)

run:
	echo "Running the package..."
	$(GO_CMD) run -ldflags="$(LDFLAGS)" $(BUILD_MAIN_FILE) $(RUN_ARGS)

run-config:
	echo "Running the package with .env..."
	$(GO_CMD) run -ldflags="$(LDFLAGS)" $(BUILD_MAIN_FILE) --config pihole_exporter.config.yml

docs:
	$(GOAUTODOC_CMD) -config ./.autodoc/config.root.yml > ./README.md
	$(GOAUTODOC_CMD) -package ./config > ./config/README.md
	$(GOAUTODOC_CMD) -package ./exporter > ./exporter/README.md
	$(GOAUTODOC_CMD) -package ./pihole > ./pihole/README.md
	$(GOAUTODOC_CMD) -package ./server > ./server/README.md

info:
	$(ECHO) "Package:    '${PACKAGE_NAME}'" > /dev/null
	$(ECHO) "Version:    '${VERSION}'" > /dev/null
	$(ECHO) "Build file: '${BUILD_MAIN_FILE}'" > /dev/null
	$(ECHO) "Go command: '${GO_CMD}'" > /dev/null
	$(ECHO) "Time:       '$(TIME)'" > /dev/null

build-all: build-linux-amd64 build-linux-arm64 build-linux-arm build-windows-amd64 build-darwin-amd64 build-darwin-arm64
