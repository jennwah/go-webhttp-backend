# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go
CMD ["./app"]
