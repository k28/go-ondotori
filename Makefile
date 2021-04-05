# メタ情報
NAME := ondotori
VERSION := $(godump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"
SRCROOT := $(NAME)
SOURCES := $(shell find $(SRCROOT) -type f -not -name '*_test.go')

export GO11MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# 開発に必要な依存をインストールする
## Setup
.PHONY: deps
devel-deps: deps
	GO11MODULE=off go get	\
	golang.org/x/lint/golint	\
	github.com/x-motemen/gobump/cmd/gobump	\
	github.com/Songmu/make2help/cmd/make2help

# テストを実行する
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries ex. make bin/ondotori
bin/%: $(SOURCES) deps
	go build -ldflags $(LDFLAGS) -o $@ $(SOURCES)

## build binary
#.PHONY: build
#build: bin/$(NAME)

## show help
#.PHONY: help
#help:
	@make2help $(MAKEFILE_LIST)

## clean
.PHONY: clean
clean:
	rm -fr bin

