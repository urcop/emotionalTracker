FROM golang:1.23.0-alpine as gogcc

ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev \
        git \
        build-base

FROM gogcc as builder

WORKDIR /build

COPY . .

RUN go mod download && go mod verify

RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main_build main.go

# production stage
FROM alpine:latest

ARG CONFIG_PATH

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev

WORKDIR /app/

COPY --from=builder /build/docs .
COPY --from=builder /build/main_build .
COPY --from=builder /build/config/ /config/


ENV CONFIG_PATH=$CONFIG_PATH


CMD ["/app/main_build"]