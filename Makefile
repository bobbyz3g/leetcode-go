GO := go

.PHONY: all
all: lint fmt test

.PHONY: test
test:
	@$(GO) test ./... -v

.PHONY: lint
lint: verify.revive
	@revive -config lint.toml ./...

.PHONY: fmt
fmt: verify.golines
	@$(GO) fmt ./...
	@golines -w .

.PHONY: verify.%
verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) install.$*; fi

.PHONY: install.revive
install.revive:
	@$(GO) get -u github.com/mgechev/revive

.PHONY: install.golines
install.golines:
	@$(GO) get -u github.com/segmentio/golines