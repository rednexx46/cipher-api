# Stage 1: build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o cipher-api main.go

# Stage 2: run the binary in a minimal image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cipher-api .

EXPOSE 8080

CMD ["./cipher-api"]