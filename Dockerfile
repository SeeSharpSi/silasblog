FROM golang:1.23.3

WORKDIR /app

COPY ../silasblog /app
RUN go build server.go

EXPOSE 9779

CMD ["go", "run", ".", "--address=0.0.0.0"]
