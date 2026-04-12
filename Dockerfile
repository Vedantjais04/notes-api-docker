# Build stage
FROM golang:1.26 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage (lightweight)
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]