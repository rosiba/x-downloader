FROM golang:alpine AS compile
LABEL authors="rosiba"

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bot-service .

FROM alpine:latest

RUN apk add --no-cache \
    python3 \
    ffmpeg \
    ca-certificates \
    curl

RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download-yt-dlp -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp

WORKDIR /root/

COPY --from=compile /app/bot-service .

EXPOSE 8000

CMD ["./bot-service"]