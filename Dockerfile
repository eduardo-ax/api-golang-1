FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN build -v -o rest_api

FROM alpine

COPY --from=builder /app/rest_api /

ENTRYPOINT [ "/rest_api" ]