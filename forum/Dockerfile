# Stage 1: Build the Go application
FROM golang:1.22.4-alpine AS build

# Install build dependencies
# this is required for the sqlite3
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary with CGO enabled
# this is required for the sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -o forum .

# Stage 2: Create a lightweight container with the binary and database
FROM alpine:3.20

# Set up environment variables
ENV APP_HOME /app

# Create the application directory
WORKDIR $APP_HOME

# Set the working directory inside the new container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/forum .

# Copy the SQLite database file (assuming it exists locally)
COPY --from=build /app/data.db ./data.db

# Expose the necessary port (change if needed)
EXPOSE 8080

# Set the entrypoint to run the binary
ENTRYPOINT ["./forum"]

