# Use an official Golang runtime as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Compile the binary inside the container
RUN go build -o main .

# Set the command to run the binary
CMD ["./main"]
