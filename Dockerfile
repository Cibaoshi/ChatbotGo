FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download || true
COPY . .
RUN go build -o /server .

FROM alpine:latest
WORKDIR /
COPY --from=builder /server /server
COPY --from=builder /app/templates /templates
ENV PORT=8080
EXPOSE 8080
CMD ["/server"]