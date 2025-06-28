# syntax=docker/dockerfile:1

# Stage 1: Build the Go binary
# Use a specific version of golang-alpine for reproducibility and small size
FROM golang:1.23-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies.
# This leverages Docker's layer caching, so dependencies are only
# re-downloaded when these files change.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code into the container.
COPY . .

# Build the Go application, creating a statically linked binary.
# CGO_ENABLED=0 is important for building a static binary that can run
# in a minimal container like Alpine.
# GOOS=linux ensures the binary is built for the Linux operating system.
# -o /server places the compiled binary in the root of the builder's filesystem.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /server .

# Stage 2: Create the final, minimal production image
# Start from a minimal base image to reduce the attack surface and image size.
FROM alpine:latest

# Set the working directory for the final image.
WORKDIR /

# Copy the necessary files from the 'builder' stage.
# We only copy the compiled binary and the static assets, not the source code
# or the Go toolchain.
COPY --from=builder /server /server
COPY data.db /data.db
COPY --from=builder /app/static /static

# Expose the port that the application listens on. Cloud Run uses this value.
# This should match the port in your server.go (default is 9779).
EXPOSE 9779

# Set the command to run when the container starts. This executes the compiled binary.
ENTRYPOINT ["/server"]
