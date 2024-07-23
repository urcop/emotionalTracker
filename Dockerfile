FROM golang:1.21.5-alpine as gogcc

ENV GOOS=linux
ENV CGO_ENABLED=1
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

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o build/app main.go

# production stage
FROM alpine:latest

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev

WORKDIR /app/

COPY --from=builder /build/docs .
COPY --from=builder /build/app .
COPY --from=builder /build/config/ /config/

CMD ["/app/app", "s"]