# Stage 1: Builder
# Use a specific Go version instead of 'latest' for reproducibility
FROM golang:1.24-alpine AS builder

# Install dependencies (if needed for building, e.g., git, private repositories)
# In most cases, this is not needed for simple Go projects
# RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first for better layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project code
COPY . .

# Build the application
# Add flags for static linking (-s, -w) and for Alpine (-extldflags '-static')
# This will allow using scratch as the base image if you want maximum minimalism
RUN CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'" -o myapp .

# Stage 2: Final image
# Use a minimal Alpine image
FROM alpine:latest

# Install necessary runtime dependencies (rarely needed for Go if build is static)
# For example, if the application requires SSL certificates or specific libraries
# RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy only the compiled executable from the builder stage
COPY --from=builder /app/myapp ./

# Set the user to run the application as (optional, but recommended for security)
# RUN adduser -u 1000 -D appuser
# USER appuser

# Set the entry point
ENTRYPOINT ["./myapp"]