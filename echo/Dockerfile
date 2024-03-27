# -----Build Stage-----
FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# アプリケーションをビルド
RUN go build -o main ./main.go
RUN go install github.com/cosmtrek/air@latest

# -----Run Stage-----
# FROM alpine:3.19.1
FROM golang:1.22.1-alpine3.19 AS runtime

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /go/bin/air /usr/local/bin/air

EXPOSE 8080

# CMD ["air", "-c", ".air.toml"]
CMD ["air"]

