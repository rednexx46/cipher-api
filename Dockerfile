# Stage 1: Build
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 🔥 Ensure static build for distroless
RUN CGO_ENABLED=0 GOOS=linux go build -o cipher-api .

# Stage 2: Run with minimal distroless image
FROM alpine:latest as run

WORKDIR /

COPY --from=builder /app/cipher-api .

ENTRYPOINT ["/cipher-api"]
