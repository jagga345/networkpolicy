# Dockerfile for m1 server
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o server ./server.go

EXPOSE 50051

CMD ["./server"]
