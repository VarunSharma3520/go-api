# =========================
# Build stage
# =========================
FROM golang:1.25 AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/go-api/main.go

# =========================
# Runtime stage
# =========================
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/app /app/app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/app"]
