FROM golang:1.19.4-alpine3.17

WORKDIR /app

RUN apk update

COPY . /app

RUN go mod tidy

CMD ["go", "run", "main.go"]