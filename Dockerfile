FROM golang:1.24-alpine AS builder
WORKDIR /migrator

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
# COPY migrations/ ./migrations
COPY config/ ./config

RUN go build -o migrator main.go

FROM alpine:latest
WORKDIR /migrations
RUN apk --no-cache add postgresql-libs
COPY --from=builder /migrator/migrator /usr/local/bin/migrator
COPY migrations/ ./migrations
ENTRYPOINT ["migrator"]
