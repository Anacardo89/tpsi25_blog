
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o tpsi25_blog cmd/main.go

FROM debian:12-slim 
WORKDIR /root/
COPY --from=builder /app/tpsi25_blog .
COPY cmd/ssl/ /root/ssl/
COPY cmd/templates/ /root/templates
COPY cmd/static/ /root/static

EXPOSE 8081
EXPOSE 8082

CMD ["./tpsi25_blog"]
