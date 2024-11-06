# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go binary
RUN go build -o binary

# Command to run the binary
CMD ["./binary"]
