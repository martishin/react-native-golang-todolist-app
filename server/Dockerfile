# Build stage
FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN ls -la
RUN GOOS=linux go build -o server ./cmd

# Final stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 3000
CMD ["/app/server"]
