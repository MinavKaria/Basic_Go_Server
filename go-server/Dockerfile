FROM golang:alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o /app/server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server /app/
COPY config.yaml /app/

EXPOSE 8080
CMD ["/app/server"]