# Dockerfile for m1 client
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o client ./client.go

CMD ["./client"]
