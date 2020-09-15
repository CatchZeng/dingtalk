SHELL := /bin/bash
BASEDIR = $(shell pwd)

export GO111MODULE=on

APP_NAME=dingtalk
APP_VERSION=1.2.0
IMAGE_NAME="catchzeng/${APP_NAME}:${APP_VERSION}"
IMAGE_LATEST="catchzeng/${APP_NAME}:latest"

all: fmt
	echo 'make all'
fmt:
	gofmt -w .
mod:
	go mod tidy
lint:
	golangci-lint run
.PHONY: test
test:
	sh scripts/test.sh
mock:
	sh scripts/mock.sh
.PHONY: build
build:
	rm dingtalk
	go build -o dingtalk cmd/main.go
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dingtalk cmd/main.go
	zip dingtalk-darwin-amd64.zip dingtalk
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dingtalk cmd/main.go
	zip dingtalk-linux-amd64.zip dingtalk
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dingtalk.exe cmd/main.go
	zip dingtalk-windows-amd64.zip dingtalk.exe
build-win32:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dingtalk.exe cmd/main.go
	zip dingtalk-windows-386.zip dingtalk.exe
build-release:
	make build-mac
	make build-linux
	make build-win
	make build-win32
build-docker:
	echo ${IMAGE_NAME}
	sh build/package/build.sh ${IMAGE_NAME}
push-docker: build-docker
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST};
	docker push ${IMAGE_NAME};
	docker push ${IMAGE_LATEST};
help:
	@echo "fmt - go format"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "build - build binary"
	@echo "build-mac - build mac binary"
	@echo "build-linux - build linux amd64 binary"
	@echo "build-win - build win amd64 binary"
	@echo "build-win32 - build win 386 binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image to docker hub"
