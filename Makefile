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
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dingtalk main.go
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dingtalk main.go
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dingtalk.exe main.go
build-win32:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dingtalk.exe main.go
utest:
	go test -coverpkg=./... -coverprofile=coverage.data ./...
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"