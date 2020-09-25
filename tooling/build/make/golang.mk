# This makefile contains common instructions to build golang artifacts

# Ensures that all variables are set
ifndef PKG
$(error "PKG is not set")
endif

ifndef VERSION
$(error "VERSION is not set")
endif

# Let's always assume someone wants to build for Linux
# If it is not the case, use 'GOOS=darwin make' instead
GOOS   ?= linux
GOARCH ?= amd64
CGO_ENABLED ?= 0
GO111MODULE ?= on
export GOOS
export GOARCH
export CGO_ENABLED
export GO111MODULE

# Basic go commands
GO       ?= go
BENCHCMP ?= benchcmp
GOTEST   ?= $(GO) test
GOGET    ?= $(GO) get -v -u
LINTER   ?= golint
GOBUILD  ?= $(GO) build -ldflags="-s -w -X main.version=$(VERSION) $(LDFLAGS)"

# Golang packages in this project
PKG_LIST = $(shell $(GO) list ./... | \
	grep -v $(PKG)/vendor \
)

# Get a list of all golang source code that could be used in building the binaries
SOURCES = $(shell find ./ \
	-name '*.go' \
	-not -name '*_test.go' \
	-not -path './vendor/*' \
)

# We will always try to compress binaries if possible
HAS_UPX_AVAILABLE := $(shell command -v upx 2> /dev/null)
ifndef HAS_UPX_AVAILABLE
$(warning "Binaries cannot be compressed, upx not available in this system")
COMPRESS_ENABLED = 0
else
COMPRESS_ENABLED ?= 1
endif

###########################
### Testing and linting ###
###########################
go-test: vendor ## Run 'go test -short' in all packages
	$(GOTEST) -short $(PKG_LIST)

go-lint: vendor ## Lint the files
	$(LINTER) -set_exit_status $(PKG_LIST)

go-install-lint: ## Install golint binary
	$(GOGET) golang.org/x/lint/golint

###################
### Development ###
###################
go-clean: ## Removes all assests created by the golang building
	rm -rf bin/* .cover coverage.html
