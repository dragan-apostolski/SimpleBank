# Use the official golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download and install dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

