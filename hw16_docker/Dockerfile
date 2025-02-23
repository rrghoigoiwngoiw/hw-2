# 1. Этап сборки (Builder)
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum, чтобы заранее скачать зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем все файлы проекта
COPY . .

# Собираем бинарник из main.go
RUN go build -o /go/bin/app ./main.go

# 2. Этап финального образа (Production)
FROM alpine:latest

WORKDIR /root/

# Копируем собранный сервер из builder-а
COPY --from=builder /go/bin/app /app

# Открываем порт 8080 (если в коде используется другой — измени)
EXPOSE 8080

# Команда запуска контейнера
CMD ["/app"]
