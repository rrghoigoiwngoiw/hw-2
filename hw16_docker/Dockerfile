FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod vendor  

COPY . .

RUN go build -o /go/bin/app . 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/app /app

EXPOSE 8080

CMD ["/app"]
