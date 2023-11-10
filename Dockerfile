# Build the application from source
FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Run the tests in the container
FROM builder AS tester
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM debian:bullseye-slim AS release

RUN apt-get update \
    && apt-get install ca-certificates -y \
    && apt-get clean

WORKDIR /

COPY --from=builder /server /server

USER nonroot:nonroot

ENTRYPOINT ["/server", "start"]
