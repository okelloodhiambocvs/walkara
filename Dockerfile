# Build stage
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o walkara ./cmd/api

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/walkara .

EXPOSE 8080

CMD ["./walkara"]