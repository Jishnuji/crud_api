# Use an official Go image as the base image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
# RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Use a minimal base image for the final container
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set working directory for runtime
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Copy config file
COPY /config/config_db.yml /app/

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["/root/main"]

# Run unit tests by default
# CMD ["go", "test", "./..."]
