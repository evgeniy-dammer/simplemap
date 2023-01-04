FROM golang:alpine AS builder
WORKDIR /app
COPY go.* ./
COPY . ./
RUN go mod download
RUN go build -o tileserver

FROM alpine
WORKDIR /app
COPY --from=builder /app/tileserver /app/tileserver
COPY tiles  /app/tiles/
CMD ["/app/tileserver"]