FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o catalog ./cmd/api

FROM alpine

WORKDIR /app

COPY --from=builder /app/catalog .

EXPOSE 8080

CMD ["./catalog"]
