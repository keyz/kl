REPO_ROOT := $(shell git rev-parse --show-toplevel)
include $(REPO_ROOT)/internal/shared.mk

.PHONY: format
format:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy
