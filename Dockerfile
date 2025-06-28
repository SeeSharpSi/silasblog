FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /server .

FROM alpine:latest

WORKDIR /

COPY --from=builder /server /server
COPY --from=builder /app/data.db /data.db
COPY --from=builder /app/static /static

EXPOSE 9779

ENTRYPOINT ["/server", "-db", "/data.db"]
