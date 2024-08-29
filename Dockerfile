# Start with an official Go runtime as a parent image
FROM golang:1.23-alpine

# Set environment variables
ENV CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /go-movies ./cmd/api

# Expose the port on which your Go API listens
EXPOSE 8080

# Set the entry point to the Go app binary
CMD ["/go-movies"]
