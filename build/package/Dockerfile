FROM golang:1.13.1-alpine as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN ls -la
RUN go build -o dingtalk cmd/main.go

FROM alpine:3.7

COPY --from=builder /build/dingtalk /usr/local/bin/dingtalk
RUN chmod +x /usr/local/bin/dingtalk
RUN dingtalk version