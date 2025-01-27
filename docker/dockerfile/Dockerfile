FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache \
    bluez \
    dbus \
    gcc \
    musl-dev

WORKDIR /go/src/github.com/Paprec/trucktrack

COPY . .

RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux\
    go build -mod=vendor -ldflags "-s -w" -o build/trucktrack-led cmd/main.go \
    && mv build/trucktrack-led /exe

FROM alpine:latest

RUN apk add --no-cache \
    bluez \
    dbus

COPY --from=builder /exe /

ENV PATH=$PATH:/usr/local/bin

EXPOSE 9090

RUN mkdir -p /var/run/dbus

ENTRYPOINT ["/exe"]