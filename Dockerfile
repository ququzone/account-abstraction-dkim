# Build the application from source
FROM golang:1.19 AS builder

WORKDIR /app

RUN --mount=type=bind,source=cmd,target=cmd \
    --mount=type=bind,source=pkg,target=pkg \
    --mount=type=bind,source=main.go,target=main.go \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    <<EOF
set -e
go mod download
CGO_ENABLED=0 GOOS=linux go build -o /bin/server
EOF

# Run the tests in the container
FROM builder AS tester
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM debian:bullseye-slim AS release

RUN apt-get update \
    && apt-get install ca-certificates -y \
    && apt-get clean

WORKDIR /

COPY --from=builder /bin/server /bin/server

CMD ["/bin/server", "start"]
