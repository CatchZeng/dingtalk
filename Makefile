SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

APP_NAME=dingtalk
APP_VERSION=1.1.0
IMAGE_NAME="catchzeng/${APP_NAME}:${APP_VERSION}"
IMAGE_LATEST="catchzeng/${APP_NAME}:latest"

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
build-docker:
	echo ${IMAGE_NAME}; \
	docker build -t ${IMAGE_NAME} -f Dockerfile .;
push-docker: build-docker
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST};
	docker push ${IMAGE_NAME};
	docker push ${IMAGE_LATEST};
utest:
	go test -coverpkg=./... -coverprofile=coverage.data ./...
help:
	@echo "fmt - go format"
	@echo "mod - go mod tidy"
	@echo "build - build binary"
	@echo "build-mac - build mac binary"
	@echo "build-linux - build linux amd64 binary"
	@echo "build-win - build win amd64 binary"
	@echo "build-win32 - build win 386 binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image to docker hub"
	@echo "utest - unit test"