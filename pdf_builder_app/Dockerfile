FROM golang:1.19.4-alpine3.17

WORKDIR /app

RUN apk update

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

COPY . /app

RUN go mod tidy

CMD ["go", "run", "main.go"]