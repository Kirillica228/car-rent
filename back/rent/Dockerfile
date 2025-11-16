FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o rent ./cmd/api

FROM alpine

WORKDIR /app

COPY --from=builder /app/rent .

EXPOSE 8080

CMD ["./rent"]
