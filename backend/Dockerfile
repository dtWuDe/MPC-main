# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

ARG CMD_PATH

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main $CMD_PATH

# Stage 2: Run the Go binary
FROM alpine:latest

# Install postgresql client (psql)
RUN apk add --no-cache postgresql-client

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy .env file
COPY .env .

# Expose port 5001 to the outside world
EXPOSE 5001

# Command to run the executable
CMD ["./main"]