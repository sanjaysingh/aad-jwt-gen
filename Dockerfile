# Start from the official Golang image to build our application
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aad-jwt-gen .

# Start a new stage from scratch for a smaller, final image
FROM alpine:latest  

# Install ca-certificates for HTTPS requests (e.g., Azure AD requests)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage and the static folder
COPY --from=builder /app/aad-jwt-gen .
COPY --from=builder /app/static ./static

# Expose port 5555 to the outside world
EXPOSE 5555

# Command to run the executable
CMD ["./aad-jwt-gen"]
