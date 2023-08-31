GO ?= go
define check_tool
	@if ! command -v $(1) >/dev/null 2>&1 ; then\
		echo "ERROR: '$(1)' binary not found. Please ensure that tool is installed or specify binary path with '$(2)' variable." && \
		exit 1; \
	fi;
endef

.PHONY: build
build: check
	go build -ldflags "-w -s" -o bin/fusion main.go

.PHONY: check
check:
	$(call check_tool,$(GO),'GO')

clean:
	@echo ":: Cleanup..." && rm -rf bin/*
