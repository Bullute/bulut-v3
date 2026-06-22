FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o biner-executable .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/biner.executable .
CMD [ "./biner-executable" ]