SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=dingtalk
APP_VERSION=1.5.0
IMAGE_NAME="catchzeng/${APP_NAME}:${APP_VERSION}"
IMAGE_LATEST="catchzeng/${APP_NAME}:latest"

all: mod fmt imports lint test
first:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fmt:
	gofmt -w .
mod:
	go mod tidy
imports:
	goimports -w .
lint:
	golangci-lint run
.PHONY: test
test:
	sh scripts/test.sh
test-sonar:
	go test -gcflags=-l -coverpkg=./... -coverprofile=coverage.data ./...
mock:
	sh scripts/mock.sh
.PHONY: build
build:
	rm -f dingtalk
	go build -o dingtalk main.go
build-darwin-amd64:
	rm -f dingtalk dingtalk-darwin-amd64.zip
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dingtalk main.go
	zip dingtalk-darwin-amd64.zip dingtalk
build-darwin-arm64:
	rm -f dingtalk dingtalk-darwin-arm64.zip
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dingtalk main.go
	zip dingtalk-darwin-arm64.zip dingtalk
build-linux:
	rm -f dingtalk dingtalk-linux-amd64.zip
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dingtalk main.go
	zip dingtalk-linux-amd64.zip dingtalk
build-win:
	rm -f dingtalk.exe dingtalk-windows-amd64.zip
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dingtalk.exe main.go
	zip dingtalk-windows-amd64.zip dingtalk.exe
build-win32:
	rm -f dingtalk.exe dingtalk-windows-386.zip
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dingtalk.exe main.go
	zip dingtalk-windows-386.zip dingtalk.exe
build-release:
	make build-darwin-amd64
	make build-darwin-arm64
	make build-linux
	make build-win
	make build-win32
	rm -f dingtalk dingtalk.exe
build-docker:
	sh build/package/build.sh ${IMAGE_NAME}
push-docker:
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST};
	docker push ${IMAGE_NAME};
	docker push ${IMAGE_LATEST};
help:
	@echo "first - first time"
	@echo "fmt - go format"
	@echo "mod - go mod tidy"
	@echo "imports - go imports"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "mock - mockgen"
	@echo "build - build binary"
	@echo "build-mac - build mac binary"
	@echo "build-linux - build linux amd64 binary"
	@echo "build-win - build win amd64 binary"
	@echo "build-win32 - build win 386 binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image to docker hub"
