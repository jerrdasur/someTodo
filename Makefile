NAME := someTodo
PATH := $(PWD)/bin:$(PATH)
LINT_VERSION := v1.45.2

PKG := `go list -mod=mod -f {{.Dir}} ./...`
MAIN := cmd/${NAME}/main.go

ifeq ($(RACE),1)
	GOFLAGS=-race
endif

.PHONY: build
build:
	@CGO_ENABLED=0 go build -mod=mod $(GOFLAGS) -o ${NAME} $(MAIN)

fmt:
	@goimports -local ${NAME} -l -w $(PKG)

lint:
	@golangci-lint run -c .golangci.yml

tool:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin ${LINT_VERSION}
	@go install golang.org/x/tools/cmd/goimports@latest

mod:
	@go mod download
	@go mod tidy

test:
	@go test -v ./...

pre-commit: mod fmt lint test
