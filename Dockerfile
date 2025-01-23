FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]