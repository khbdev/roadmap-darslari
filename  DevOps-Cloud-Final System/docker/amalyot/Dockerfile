
FROM golang:1.22-alpine AS builder

WORKDIR /src
COPY main.go .


RUN go build -o app main.go



FROM alpine:latest
COPY --from=builder /src/app /app
CMD ["/app"]
