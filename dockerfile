FROM golang:1.17-alpine as builder

WORKDIR /go/src/core

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .

FROM mariadb:latest

WORKDIR /go/src/core

COPY --from=builder /go/src/core/main .
COPY --from=builder /go/src/core/.env .

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["mariadbd"]