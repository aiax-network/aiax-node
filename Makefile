#!/usr/bin/make -f

#VERSION ?= $(shell echo $(shell git describe --tags `git rev-list --tags="v*" --max-count=1`) | sed 's/^v//')
VERSION ?= v0.0.1
TMVERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
BINDIR ?= $(GOPATH)/bin
AIAX_BINARY = aiaxd
AIAX_DIR = aiax
BUILDDIR ?= $(CURDIR)/build
PROJECT := aiax

export GO111MODULE = on

default_target: all


build_tags = netgo ledger gcc
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=aiax \
			-X github.com/cosmos/cosmos-sdk/version.AppName=$(AIAX_BINARY) \
			-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
			-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
			-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TMVERSION)

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
	ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  #BUILD_FLAGS += -trimpath
endif

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/
$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

distclean: clean tools-clean

clean:
	rm -rf \
    $(BUILDDIR)/ \
    artifacts/ \
    tmp-swagger-gen/

all: build

.PHONY: default_target all build distclean clean
