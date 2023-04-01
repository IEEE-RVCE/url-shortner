# Base image
FROM golang:1.17-alpine

# Set working directory
WORKDIR /app

# Copy source code to container
COPY . .

# Install dependencies
RUN go mod download

# Build the binary
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
