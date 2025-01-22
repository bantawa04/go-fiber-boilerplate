FROM golang:1.23-alpine

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/air-verse/air@latest

COPY . /app

WORKDIR /app

RUN go mod tidy