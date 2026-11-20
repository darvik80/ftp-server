# =============================================================================
#  ftp-server-go — Makefile
#
#  Кроссплатформенный: работает из cmd.exe / PowerShell (mingw/mainline make)
#  и из sh/bash (Linux, macOS, Git Bash, WSL).
#
#  Требования: Go >= 1.27, GNU Make >= 4.0
#  Windows:    scoop install make   (или choco install make)
#
#  Быстрый старт:   make            # список целей
#                   make test       # тесты
#                   make build      # build\ftp-server-go.exe
# =============================================================================

# ------------------------------ Project --------------------------------------
MODULE      := ftp-server-go
BINARY      := ftp-server-go
MAIN_PKG    := .
CMD_PKG     := ./cmd/...
BUILD_DIR   := build
COVER_FILE  := coverage.out
COVER_HTML  := coverage.html

VERSION     ?= $(shell git describe --tags --always --dirty 2>$(NULL))
GIT_COMMIT  ?= $(shell git rev-parse --short HEAD 2>$(NULL))

# ------------------------------ Platform -------------------------------------
ifeq ($(OS),Windows_NT)
    SHELL        := cmd.exe
    .SHELLFLAGS  := /C
    EXE          := .exe
    NULL         := nul
    MKDIR        := if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
    RMDIR        := if exist "$(BUILD_DIR)" rmdir /S /Q "$(BUILD_DIR)"
    ENV_STATIC   := set CGO_ENABLED=0&&
    GOENV        =  set GOOS=$(word 1,$(subst _, ,$1))&& set GOARCH=$(word 2,$(subst _, ,$1))&&
    HELP_CMD     = powershell -NoProfile -Command "Select-String -Path '$(firstword $(MAKEFILE_LIST))' -Pattern '^([A-Za-z0-9_.-]+):.*##\s+(.+)\$$' | ForEach-Object { $_.Matches[0].Groups[1].Value.PadRight(16) + $_.Matches[0].Groups[2].Value }"
else
    SHELL        := /usr/bin/env bash
    .SHELLFLAGS  := -e -o pipefail -c
    EXE          :=
    NULL         := /dev/null
    MKDIR        := mkdir -p $(BUILD_DIR)
    RMDIR        := rm -rf $(BUILD_DIR)
    ENV_STATIC   := CGO_ENABLED=0
    GOENV        =  GOOS=$(word 1,$(subst _, ,$1)) GOARCH=$(word 2,$(subst _, ,$1))
    HELP_CMD      = awk 'BEGIN {FS = ":.*?## "} /^[A-Za-z0-9_.-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(firstword $(MAKEFILE_LIST))
endif

BUILD_DATE  := $(shell git log -1 --format=%cI 2>$(NULL))

# ldflags: обрезка отладочной инфы + инъекция версии
LDFLAGS     := -s -w
LDFLAGS     += -X main.version=$(if $(VERSION),$(VERSION),dev)
LDFLAGS     += -X main.commit=$(if $(GIT_COMMIT),$(GIT_COMMIT),unknown)
LDFLAGS     += -X main.date=$(if $(BUILD_DATE),$(BUILD_DATE),unknown)

GO          := go
GOBUILD     := $(GO) build -trimpath -ldflags "$(LDFLAGS)"
GOTEST      := $(GO) test -count=1
GOFLAGS     ?=
LDFLAGS_PKG := ./...

# Целевые платформы для make build-all:  os_arch
PLATFORMS   := linux_amd64 linux_arm64 windows_amd64 darwin_amd64 darwin_arm64
RELEASE_BINS := $(foreach p,$(PLATFORMS),$(BUILD_DIR)/$(BINARY)-$(p)$(if $(findstring windows,$(p)),.exe))

.DEFAULT_GOAL := help
.PHONY: help deps tidy verify fmt fmt-check imports vet lint staticcheck \
        test test-v test-short test-race cover cover-html \
        build build-all run install clean version check ci tools

# ============================== Help =========================================
help: ## Show this help
	@$(HELP_CMD)

version: ## Print version / build info
	@$(GO) version
	@echo version = $(if $(VERSION),$(VERSION),dev)
	@echo commit  = $(if $(GIT_COMMIT),$(GIT_COMMIT),unknown)
	@echo module  = $(MODULE)

# ============================== Dependencies =================================
deps: ## Download modules
	@$(GO) mod download

tidy: ## go mod tidy (+ vendor prune)
	@$(GO) mod tidy
	@$(GO) mod verify

verify: ## Verify modules checksum
	@$(GO) mod verify

tools: ## Install dev tools (goimports, golangci-lint, staticcheck)
	@$(GO) install golang.org/x/tools/cmd/goimports@latest
	@$(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	@$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	@$(GO) install honnef.co/go/tools/cmd/staticcheck@latest

# ============================== Quality ======================================
fmt: ## gofmt -s -w
	@gofmt -s -w .

fmt-check: ## Fail if sources are not gofmt-ed
	@$(if $(shell gofmt -s -l .),echo "not formatted -> run 'make fmt'" && exit 1,echo "format: OK")

imports: ## goimports (groups std / external)
	@goimports -local $(MODULE)/ -w .

vet: ## go vet
	@$(GO) vet ./...

lint: ## golangci-lint run (skipped if not installed)
ifeq ($(strip $(shell where golangci-lint 2>$(NULL))),)
	@echo "golangci-lint NOT FOUND -> skip (make tools)"
else
	@golangci-lint run ./...
endif

staticcheck: ## staticcheck (skipped if not installed)
	@$(GO) run honnef.co/go/tools/cmd/staticcheck@latest ./...

check: fmt-check vet test ## Full local gate: format + vet + tests

ci: fmt-check vet staticcheck lint test-race cover ## CI pipeline

# ============================== Tests ========================================
test: ## Run unit tests
	@$(GOTEST) $(LDFLAGS_PKG)

test-v: ## Tests, verbose
	@$(GOTEST) -v $(LDFLAGS_PKG)

test-short: ## Skip long/integration tests (-short)
	@$(GOTEST) -short $(LDFLAGS_PKG)

test-race: ## Tests with race detector (needs CGO/gcc on Windows)
	@$(GOTEST) -race $(LDFLAGS_PKG)

cover: ## Coverage report (func totals)
	@$(GOTEST) -coverprofile=$(COVER_FILE) -covermode=atomic $(LDFLAGS_PKG)
	@$(GO) tool cover -func=$(COVER_FILE) | findstr /I "total:" 2>$(NULL) || $(GO) tool cover -func=$(COVER_FILE) | tail -n 1

cover-html: ## coverage.html
	@$(GOTEST) -coverprofile=$(COVER_FILE) -covermode=atomic $(LDFLAGS_PKG)
	@$(GO) tool cover -html=$(COVER_FILE) -o $(COVER_HTML)
	@echo "written: $(BUILD_DIR)/$(COVER_HTML)"

# ============================== Build ========================================
build: ## Build binary -> $(BUILD_DIR)/$(BINARY)$(EXE)
	@$(MKDIR)
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY)$(EXE) $(MAIN_PKG)
	@echo "built: $(BUILD_DIR)/$(BINARY)$(EXE)"

static: ## Build CGO_ENABLED=0 for current platform
	@$(MKDIR)
	@$(ENV_STATIC) $(GOBUILD) -o $(BUILD_DIR)/$(BINARY)-static$(EXE) $(MAIN_PKG)

build-all: $(RELEASE_BINS) ## Cross-compile for $(PLATFORMS)

# Pattern rule: build/ftp-server-linux_amd64 -> GOOS=linux GOARCH=amd64
$(BUILD_DIR)/$(BINARY)-%:
	@$(MKDIR)
	@$(ENV_STATIC) $(call GOENV,$*) $(GOBUILD) -o $@ $(MAIN_PKG)

run: ## Run app: make run ARGS="-addr :2121"
	@$(GO) run $(MAIN_PKG) $(ARGS)

install: ## go install binary
	@$(GO) install $(MAIN_PKG)

# ============================== Clean ========================================
clean: ## Remove build artifacts
	@$(RMDIR)
	@$(if $(shell cmd /c if exist $(COVER_FILE) echo y 2>$(NULL)),del /Q $(COVER_FILE),)
	@$(if $(shell test -f $(COVER_FILE) && echo y),rm -f $(COVER_FILE),)
	@$(GO) clean -cache -testcache 2>$(NULL) || exit 0
	@echo "cleaned"