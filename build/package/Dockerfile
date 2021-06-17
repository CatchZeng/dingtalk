FROM golang:1.16.4-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy && -o dingtalk cmd/main.go

FROM alpine:3.7

COPY --from=builder /build/dingtalk /usr/local/bin/dingtalk
RUN chmod +x /usr/local/bin/dingtalk
