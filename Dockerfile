# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git jika ada go.mod yang butuh ambil module
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary statis dan kecil
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" .

# Stage 2: Final image (ringan)
FROM alpine:latest

WORKDIR /root/

# Copy hasil build dari stage builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
