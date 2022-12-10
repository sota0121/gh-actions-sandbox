.PHONY: build clean test run

GO ?= go
MAIN_FILE ?= cmd/main.go
OUT_DIR ?= out
BINARY_NAME ?= sandbox
BIN_DIR ?= $(OUT_DIR)/bin
BIN_PATH ?= $(BIN_DIR)/$(BINARY_NAME)
COVERAGE_FILE ?= $(OUT_DIR)/coverage.out

build:
	$(GO) build -o $(BIN_PATH) $(MAIN_FILE)

clean:
	rm -rf $(OUT_DIR)

test:
	$(GO) test -coverprofile=$(COVERAGE_FILE) ./...

run:
	$(GO) run $(MAIN_FILE)
