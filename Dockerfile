# builder stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /app/bin/svc

# app stage
FROM alpine:latest AS production

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/bin/svc /app/bin/svc
EXPOSE 8080

ENTRYPOINT ["/app/bin/svc"]
