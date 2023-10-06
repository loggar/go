REPO_NAME := app-1
VERSION_VAR := main.version
GIT_COMMIT_VAR := main.gitCommit
REPO_VERSION := $$(git describe --tags 2>/dev/null || echo nil)
GIT_HASH := $$(git rev-parse --short HEAD)
GOBUILD_VERSION_ARGS := -ldflags "-s -X '$(VERSION_VAR)=$(REPO_VERSION)' -X '$(GIT_COMMIT_VAR)=$(GIT_HASH)'"
BINARY_NAME := $(REPO_NAME)
BINARY_NAME_WORKER := $(REPO_NAME)-worker
BINARY_OUT_DIR := ./_target/cmd
COVER_DIR := ./_target/test
COVER_OUT := coverage.out

build-all: depend depend-tidy cover build

depend:
	go mod download

depend-tidy:
	go mod tidy -v

build:
	go build $(GOBUILD_VERSION_ARGS) -o $(BINARY_OUT_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
	go build $(GOBUILD_VERSION_ARGS) -o $(BINARY_OUT_DIR)/$(BINARY_NAME_WORKER) ./cmd/$(BINARY_NAME_WORKER)

clean:
	rm -f $(COVER_DIR)/$(COVER_OUT)
	rm -f $(BINARY_OUT_DIR)/$(BINARY_NAME)
	rm -f $(BINARY_OUT_DIR)/$(BINARY_NAME_WORKER)

test:
	go test -vet=off ./...

test-race:
	go test -race -vet=off ./...

bench:
	go test -bench=. -run=XXX ./...

bench-race:
	go test -race -bench=. -run=XXX ./...

cover:
	go test -cover ./... -coverprofile=$(COVER_DIR)/$(COVER_OUT)

cover-report:
	$(MAKE) cover || true
	go tool cover -func=$(COVER_DIR)/$(COVER_OUT)
	go tool cover -html=$(COVER_DIR)/$(COVER_OUT)

# None of the Make tasks generates files with the name of the task, so all must be declared as 'PHONY'
.PHONY: bench bench-race build build-all cover cover-report depend test test-race
