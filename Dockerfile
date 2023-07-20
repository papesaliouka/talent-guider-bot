# Build stage
FROM golang AS builder

LABEL maintainer="Z01-Student"

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy only the go.mod and go.sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download Go modules dependencies (this layer will be cached if go.mod and go.sum haven't changed)
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Production stage
FROM alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /go/src/app/app .

# Expose the necessary ports
EXPOSE 80
EXPOSE 443

# Run the Go application
CMD ["./app"]

