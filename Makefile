#############################
# Global vars
#############################
PROJECT_NAME := $(shell basename $(shell pwd))

SRCDIR       ?= .
GO            = go
TEST_RUNNER  ?= gotestsum

default: build

clean:
	@echo "=== $(PROJECT_NAME) === [ clean            ]: cleaning go..."
	@$(GO) clean -modcache -testcache -cache

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-newrelic ~/.tflint.d/plugins

# Import fragments
include build/test.mk
include build/lint.mk
include build/deps.mk