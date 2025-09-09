FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth ./cmd/api

FROM alpine

WORKDIR /app

COPY --from=builder /app/auth .

RUN chmod +x /app/auth

EXPOSE 8080

CMD ["./auth"]
