VERSION := "v1.0.0"
# VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')
DIRTY := $(shell git status --porcelain | wc -l | xargs)

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

ldflags = -X github.com/strangelove-ventures/bech32cli/cmd.Version=$(VERSION) \
					-X github.com/strangelove-ventures/bech32cli/cmd.Commit=$(COMMIT) \
					-X github.com/strangelove-ventures/bech32cli/cmd.Dirty=$(DIRTY)

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

all: install build-static-amd64 build-static-arm64

build: go.sum
	@echo "building bech32 binary..."
	@go build -mod=readonly -o build/bech32 -ldflags '$(ldflags)' .

install: go.sum
	@echo "installing bech32 binary..."
	@go build -mod=readonly -o $(GOBIN)/bech32 -ldflags '$(ldflags)' .

build-static: build-static-amd64 build-static-arm64

build-static-amd64:
	@echo "building bech32 amd64 static binary..."
	@GOOS=linux GOARCH=amd64 go build -o build/bech32-amd64 -a -tags netgo -ldflags '$(ldflags) -extldflags "-static"' .

build-static-arm64:
	@echo "building bech32 arm64 static binary..."
	@GOOS=linux GOARCH=arm64 go build -o build/bech32-arm64 -a -tags netgo -ldflags '$(ldflags) -extldflags "-static"' .

.PHONY: all build build-static-amd64 build-static-arm64