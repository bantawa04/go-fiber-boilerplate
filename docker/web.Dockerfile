FROM golang:1.23-alpine

RUN apk add --no-cache make build-base

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/air-verse/air@latest

COPY . /app

WORKDIR /app

RUN go mod tidy