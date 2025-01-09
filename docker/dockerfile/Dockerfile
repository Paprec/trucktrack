FROM golang:1.23-alpine AS builder

WORKDIR /go/src/github.com/Paprec/trucktrack
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 \
    go build -mod=vendor -ldflags "-s -w" -o build/tructtrack-led cmd/main.go \
    && mv build/tructtrack-led /exe 

FROM scratch
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]