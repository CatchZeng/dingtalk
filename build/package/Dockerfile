FROM golang:1.18.1-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy && go build -o dingtalk main.go


FROM alpine:3.7

COPY --from=builder /build/dingtalk /usr/local/bin/dingtalk
RUN chmod +x /usr/local/bin/dingtalk
