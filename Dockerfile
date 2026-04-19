FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]