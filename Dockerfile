FROM golang:1.24.0-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN go build -o rarible-api ./cmd/api/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/rarible-api ./rarible-api
COPY .env .env
EXPOSE 8081
CMD ["./rarible-api"]