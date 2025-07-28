# 1. Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Install git if your go.mod needs it (e.g., private repos)
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the app using the main in ./cmd
RUN go build -o app ./cmd

# 2. Minimal run image
FROM alpine:latest

WORKDIR /root/

ENV PORT=8080

# Copy the binary from builder stage
COPY --from=builder /app/app .

CMD ["./app"]
