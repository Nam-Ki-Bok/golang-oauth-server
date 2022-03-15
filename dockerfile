FROM golang:1.17 as builder

WORKDIR /go/src/core

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .

FROM mariadb:latest

WORKDIR /go/src/core

COPY --from=builder /go/src/core/main .
COPY --from=builder /go/src/core/.env .

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["mariadbd"]