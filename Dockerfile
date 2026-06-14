# ==========================
# Builder Stage
# ==========================
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o tasks .

# ==========================
# Runtime Stage
# ==========================
FROM alpine:3.19

RUN adduser -D appuser

WORKDIR /home/appuser

COPY --from=builder /app/tasks /usr/local/bin/tasks

USER appuser

ENTRYPOINT ["tasks"]
CMD ["--help"]
