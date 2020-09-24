GO   := go
MAIN_PATH := cmd/api/main.go

#-------------------------
# Build artefacts
#-------------------------
.PHONY: build build.deliverymuch-challenge

## Build all binaries
build:
	$(GO) build -o out/bin/deliverymuch-challenge cmd/api/main.go

run:
	$(GO) run $(MAIN_PATH)