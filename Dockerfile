# ------ BUILDER BLOCK ------ #
FROM golang:1.17.6 as builder

# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

# Add a work directory
WORKDIR /go-demo-api

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build go-demo-api
RUN go build main.go

# ------ PRODUCTION BLOCK ------ #
FROM alpine:latest as production

# Add certificates
RUN apk add --no-cache ca-certificates

# Copy .env from builder
COPY --from=builder /go-demo-api/.env .

# Copy built binary from builder
COPY --from=builder /go-demo-api/main .

# Exec built binary
CMD ./main