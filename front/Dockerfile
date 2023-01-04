FROM golang:alpine AS builder
WORKDIR /app
COPY go.* ./
COPY . ./
RUN go mod download
RUN go build -o front

FROM alpine
WORKDIR /app
COPY --from=builder /app/front /app/front
COPY html  /app/html/
CMD ["/app/front"]