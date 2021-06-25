FROM golang:alpine AS builder
ADD ./ /data/app/
WORKDIR /data/app/

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

EXPOSE 8088

RUN CGO_ENABLED=0 go build main.go

ENTRYPOINT ["./main"]