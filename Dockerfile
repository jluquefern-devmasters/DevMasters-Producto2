FROM golang:1.22 AS builder

WORKDIR /app

COPY DevMasters-Producto1.go .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o webapp DevMasters-Producto1.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/webapp .
COPY static ./static

EXPOSE 8080

CMD ["./webapp"]
