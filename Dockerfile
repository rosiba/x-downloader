FROM golang:alpine AS builder
LABEL authors="rosiba"

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bot-service .

WORKDIR /root/

COPY --from=builder /app/bot-service .

EXPOSE 8000

CMD ["./bot-service"]