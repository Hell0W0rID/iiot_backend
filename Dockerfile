# Multi-stage Dockerfile for IIOTBackend

# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata postgresql-client redis

# Set working directory
WORKDIR /app

# Copy go mod files first
COPY go.mod go.sum ./

# Update Go modules to latest versions compatible with Go 1.21
RUN go mod download && go mod tidy

# Copy source code
COPY . .

# Build the application with Go 1.21 optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o iiot-backend .

# Final stage
FROM alpine:3.18

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata curl postgresql-client

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/iiot-backend .

# Copy migrations and scripts
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/scripts ./scripts

# Make scripts executable
RUN chmod +x ./scripts/*.sh

# Create logs directory
RUN mkdir -p /app/logs && \
    chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/health || exit 1

# Run the application with service-specific startup
CMD ["./scripts/docker-start.sh"]
