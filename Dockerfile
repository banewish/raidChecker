# Dockerfile for Go app
FROM golang:1.24.1

WORKDIR /app


COPY . .

RUN go mod download
RUN go build -o app .

CMD ["./app"]