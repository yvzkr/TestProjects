# Multi-stage build for Zebrash
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o zebrash main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/zebrash .

# Create directories for input/output
RUN mkdir -p /app/input /app/output

# Copy test data
COPY --from=builder /app/testdata /app/testdata

# Expose port for web interface
EXPOSE 8080

# Default command
CMD ["./zebrash"] 