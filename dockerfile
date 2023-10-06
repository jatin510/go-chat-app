# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o main .

# Expose a port (if your Go application listens on a specific port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
