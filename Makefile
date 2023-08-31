GO ?= go
GOHOSTOS:=$(shell go env GOHOSTOS)
define check_tool
	@if ! command -v $(1) >/dev/null 2>&1 ; then\
		echo "ERROR: '$(1)' binary not found. Please ensure that tool is installed or specify binary path with '$(2)' variable." && \
		exit 1; \
	fi;
endef

.PHONY: build
build: check
	CGO_ENABLED=0 GOOS=$(GOHOSTOS)  GOARCH=amd64 go build -ldflags "-w -s" -o bin/$(GOHOSTOS)/fusion main.go

.PHONY: check
check:
	$(call check_tool,$(GO),'GO')
.PHONY: clean
clean:
	@echo ":: Cleanup..." && rm -rf bin/*
