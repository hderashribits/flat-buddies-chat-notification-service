# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o notification-service ./cmd/main.go

# Runtime Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/notification-service .

EXPOSE 8081

CMD ["./notification-service"]
