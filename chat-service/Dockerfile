# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o chat-service ./cmd/main.go

# Runtime Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/chat-service .

EXPOSE 8080

CMD ["./chat-service"]
