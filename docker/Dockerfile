# First stage: build the executable.
FROM golang:1.22-alpine AS builder

# Install (required for fetching Go modules) and the build dependencies for CGo.
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go mod tidy

# Since we're using Alpine, which uses musl libc, CGo should be enabled by default.
# However, it's good practice to explicitly enable it when you know you'll need it.
ENV CGO_ENABLED=1

# Build the Go app
RUN go build -o notes

# Second stage: copy the executable to a fresh image and run it.
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previons stage
COPY --from=builder /app/notes /app/notes

# Copy configs from local
COPY configs/ /app/configs/

# Command to run the executable
CMD ["/bin/sh", "-c", "./notes migrate --config=/app/configs/config.yaml && ./notes run --config=/app/configs/config.yaml"]
