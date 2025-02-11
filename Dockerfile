# Primeiro estágio (builder)
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
 
RUN go build -v -o api-golang-1 ./cmd/api-golang-1

FROM alpine

WORKDIR /app

COPY --from=builder /app/api-golang-1 /app/api-golang-1

ENTRYPOINT ["./api-golang-1"]