#FROM golang:1.16-alpine AS builder
FROM golang:alpine

RUN apk update && apk add --no-cache git curl

# Move to working directory (/app).
WORKDIR /app

# Copy the code into the container.
COPY . .

# Download all the dependencies that are required in your source files and update go.mod file with that dependency.
# Remove all dependencies from the go.mod file which are not required in the source files.
RUN go mod tidy

# Build the application server.
RUN go build -o binary .

HEALTHCHECK --interval=5s --timeout=3s CMD ["sh", "-c", "curl http://127.0.0.1:${SERVER_PORT}/health || exit 1"]

# Command to run when starting the container.
ENTRYPOINT ["/app/binary"]
