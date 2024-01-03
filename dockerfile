FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/main

EXPOSE 8080

CMD ["./bin/main"]