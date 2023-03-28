FROM golang:latest

# Set the working directory
WORKDIR /app

COPY go.mod ./
RUN go mod download
# Copy the Go source files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
