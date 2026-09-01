# Use latest Golang base image
FROM golang:latest AS builder

# Set working directory in the container
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Install dependency
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Start new build image
FROM alpine:latest

# Install certificates
RUN apk --no-cache add ca-certificates

# Set work directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
ENTRYPOINT ["./main"]