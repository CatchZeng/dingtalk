SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

all: fmt
	echo 'make all'
fmt:
	gofmt -w .
mod:
	go mod tidy
build:
	go build -o dingtalk main.go
utest:
	go test -coverpkg=./... -coverprofile=coverage.data ./...
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"