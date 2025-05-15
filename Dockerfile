# Use the official Golang image as the base image
FROM golang:1.24.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/app/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]