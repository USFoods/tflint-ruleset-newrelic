#
# Makefile fragment for Linting
#

GO               ?= go
GOFMT            ?= gofmt
GOIMPORTS        ?= goimports
GOLINTER         ?= golangci-lint
GO_MOD_OUTDATED  ?= go-mod-outdated
MISSPELL         ?= misspell

COMMIT_LINT_CMD   ?= go-gitlint
COMMIT_LINT_REGEX ?= "(chore|docs|feat|fix|refactor|tests?)(\([^\)]+\))?: .*"
COMMIT_LINT_START ?= "2020-06-24"


SRCDIR          ?= .
GO_PKGS         ?= $(shell ${GO} list ./... | grep -v -e "/example")
FILES           ?= $(shell find ${SRCDIR} -type f | grep -v -e '.git/' -e '/vendor/' -e 'go.sum')
GO_FILES        ?= $(shell find $(SRCDIR) -type f -name "*.go" | grep -v -e ".git/" -e '/example/')
PROJECT_MODULE  ?= $(shell $(GO) list -m)
GO_PKGS         ?= $(shell $(GO) list ./...)

lint: deps gofmt golangci goimports
lint-fix: deps gofmt-fix

gofmt: deps
	@echo "=== $(PROJECT_NAME) === [ gofmt            ]: Checking file format with $(GOFMT)..."
	@$(GOFMT) -e -l -s -d $(GO_FILES)

gofmt-fix: deps
	@echo "=== $(PROJECT_NAME) === [ gofmt-fix        ]: Fixing file format with $(GOFMT)..."
	@$(GOFMT) -e -l -s -w $(GO_FILES)

goimports: deps
	@echo "=== $(PROJECT_NAME) === [ goimports        ]: Checking imports with $(GOIMPORTS)..."
	@$(GOIMPORTS) -l -w -local $(PROJECT_MODULE) $(GO_FILES)

golangci: deps
	@echo "=== $(PROJECT_NAME) === [ golangci-lint    ]: Linting using $(GOLINTER) ($(COMMIT_LINT_CMD))..."
	@$(GOLINTER) run

.PHONY: lint gofmt gofmt-fix lint-fix golangci goimports