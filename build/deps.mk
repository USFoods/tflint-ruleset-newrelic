#
# Makefile fragment for installing deps
#

GO               ?= go
GOFMT            ?= gofmt
VENDOR_CMD       ?= ${GO} mod tidy

# Go file to track tool deps with go modules
TOOL_DIR     ?= tools
TOOL_CONFIG  ?= $(TOOL_DIR)/tools.go

GOTOOLS ?= $(shell cd $(TOOL_DIR) && go list -f '{{ .Imports }}' -tags tools |tr -d '[]')

tools:
	@echo "=== $(PROJECT_NAME) === [ tools            ]: Installing tools required by the project..."
	@cd $(TOOL_DIR) && $(VENDOR_CMD)
	@cd $(TOOL_DIR) && $(GO) install $(GOTOOLS)

deps: tools deps-only

deps-only:
	@echo "=== $(PROJECT_NAME) === [ deps             ]: Installing package dependencies..."
	@$(VENDOR_CMD)

.PHONY: deps deps-only tools