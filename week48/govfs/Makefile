# vim: set noexpandtab fo-=t:
# https://www.gnu.org/software/make/manual/make.html
SHELL=bash
.SHELLFLAGS=-ec -o pipefail
current_makefile:=$(lastword $(MAKEFILE_LIST))
current_makefile_dir:=$(dir $(abspath $(current_makefile)))

tools_dir=var/tools/
export PATH:=$(current_makefile_dir)$(tools_dir):$(PATH)
XDG_CACHE_HOME?=$(HOME)/.cache

golangci_lint_bin=$(tools_dir)golangci-lint

.PHONY: all
all: ## do everything (default target)

.PHONY: validate-static
validate-static: ## run static validation
	$(golangci_lint_bin) run $(if $(filter all commands,$(VERBOSE)),-v) ./...

.PHONY: validate-fix
validate-fix: ## fix auto-fixable validation errors
	$(golangci_lint_bin) run $(if $(filter all commands,$(VERBOSE)),-v) --fix ./...

.PHONY: test
test:  ## run tests
test coverage.out:
	go test -cover -race \
		-coverprofile=coverage.out -covermode=atomic \
		$(if $(filter all commands,$(VERBOSE)),-v) \
		$(if $(gotest_files),$(gotest_files),./...) \
		$(gotest_args) \

.PHONY: view-coverage
view-coverage: coverage.out ## view coverage
	go tool cover -html=coverage.out

.PHONY: validate
validate: validate-static test ## validate everything

.PHONY: build
all: build
build: validate
	go build ./...

.PHONY: install
install:  ## install everything
	go install ./...


.PHONY: run
run:
	go run ./cmd/govfsdemo/

.PHONY: clean-markers
clean: clean-markers
clean-markers:
	@{ test -d $(*) && { set -x; rm -vr test/data/*.marker-*; set +x; } } || echo "test/data/*.marker-* does not exist ... nothing to clean"

########################################################################
# toolchain
########################################################################

golangci_lint_version=1.43.0
golangci_lint_qualified=golangci-lint-$(golangci_lint_version)-linux-amd64.tar.gz
$(golangci_lint_bin): | $(tools_dir) $(XDG_CACHE_HOME)/
	cd $(XDG_CACHE_HOME) \
		&& wget -c -O $(golangci_lint_qualified) \
			https://github.com/golangci/golangci-lint/releases/download/v$(golangci_lint_version)/$(golangci_lint_qualified) \
		&& sha256sum -c $(abspath $(golangci_lint_qualified).sha256sum)
	tar -zxvf $(XDG_CACHE_HOME)/$(golangci_lint_qualified) --strip-components=1 -C $(tools_dir)

.PHONY: toolchain
toolchain: ## install toolchain
toolchain: $(golangci_lint_bin)

.PHONY: distclean
distclean: ## restore repo to pristine state
distclean: clean-$(tools_dir)/

########################################################################
# utility
########################################################################

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean-%/
clean-%/:
	@{ test -d $(*) && { set -x; rm -vr $(*); set +x; } } || echo "directory $(*) does not exist ... nothing to clean"

.PRECIOUS: %/
%/:
	mkdir -vp $(@)
