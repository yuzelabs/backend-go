FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

RUN go build -o bin/main

EXPOSE 8080

CMD ["air -c .air.toml --build.cmd 'go build -o bin/main' --build.bin './bin/main'"]