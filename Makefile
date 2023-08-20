REPO_NAME := app-1
VERSION_VAR := main.version
GIT_COMMIT_VAR := main.gitCommit
REPO_VERSION := $$(git describe --tags 2>/dev/null || echo nil)
GIT_HASH := $$(git rev-parse --short HEAD)
GOBUILD_VERSION_ARGS := -ldflags "-s -X '$(VERSION_VAR)=$(REPO_VERSION)' -X '$(GIT_COMMIT_VAR)=$(GIT_HASH)'"
BINARY_NAME := $(REPO_NAME)
BINARY_NAME_WORKER := $(REPO_NAME)-worker
BINARY_OUT_DIR := ./_out/cmd

build:
	go build $(GOBUILD_VERSION_ARGS) -o ${BINARY_OUT_DIR}/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
	go build $(GOBUILD_VERSION_ARGS) -o ${BINARY_OUT_DIR}/$(BINARY_NAME_WORKER) ./cmd/$(BINARY_NAME_WORKER)
