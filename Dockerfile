# ──────────────────────────────────────────────
# Stage: development (hot-reload with Air)
# ──────────────────────────────────────────────
FROM golang:1.26-alpine AS development

# Install Air for hot-reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Cache dependencies first (only re-downloaded when go.mod/go.sum change)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Air watches .go files and rebuilds automatically
CMD ["air", "-c", ".air.toml"]


# ──────────────────────────────────────────────
# Stage: builder (for production binary)
# ──────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./tmp/main ./cmd


# ──────────────────────────────────────────────
# Stage: production (lean final image)
# ──────────────────────────────────────────────
FROM alpine:3.23 AS production

WORKDIR /app

COPY --from=builder /app/tmp/main ./main
COPY --from=builder /app/.env ./.env

EXPOSE 5000

CMD ["./main"]