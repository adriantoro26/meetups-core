# syntax=docker/dockerfile:1

## Building stage
FROM golang:1.20 as build

# Define working directory
WORKDIR /app

# Copy files used for compiling the project
COPY . .

# Download dependencies from go.mod and go.sum
RUN go mod download

# Build source with CGO_ENABLED=0 to enable statically linked binaries to make the application more portable
RUN CGO_ENABLED=0 go build -o bin src/app.go

## Deployment stage
FROM gcr.io/distroless/static-debian11

# Define working directory
WORKDIR /root/

# Copy binary file from building stage
COPY --from=build /app/bin ./

# Export 8080 port
EXPOSE 8080

# Execute binary file to start the application
ENTRYPOINT ["./bin"]