FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o crm ./cmd/api

FROM alpine

WORKDIR /app

COPY --from=builder /app/crm .

EXPOSE 8080

CMD ["./crm"]
