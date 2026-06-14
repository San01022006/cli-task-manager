# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/tasks .

# Final minimal image
FROM alpine:3.19

LABEL org.opencontainers.image.title="CLI Task Manager"
LABEL org.opencontainers.image.description="A fast terminal task manager written in Go"
LABEL org.opencontainers.image.source="https://github.com/San01022006/cli-task-manager"

COPY --from=builder /bin/tasks /usr/local/bin/tasks

# Create data directory
RUN mkdir -p /root

ENTRYPOINT ["tasks"]
CMD ["--help"]
