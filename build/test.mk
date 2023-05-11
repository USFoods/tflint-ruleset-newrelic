#
# Makefile fragment for Testing
#

GO           ?= go
TEST_RUNNER  ?= gotestsum

COVERMODE    ?= atomic
SRCDIR       ?= .
GO_PKGS      ?= $(shell $(GO) list ./...)

PROJECT_MODULE ?= $(shell $(GO) list -m)

test: test-unit

test-unit: tools
	@echo "=== $(PROJECT_NAME) === [ test-unit        ]: running unit tests..."
	@$(TEST_RUNNER) -f testname --packages "$(GO_PKGS)" -- -v -parallel 10 -covermode=$(COVERMODE) 

.PHONY: test test-unit