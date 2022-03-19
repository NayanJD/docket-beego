FROM golang:1.17

ENV RUN_MODE=prod
ENV PORT=8080
ENV DB_CONN_STR=postgres://beego:password@localhost:5432/docket_local2?sslmode=disable

WORKDIR /go/src

COPY . /go/src

RUN go mod download

RUN go build

CMD ["/go/src/docket-beego"]

# CMD ["sleep", "infinity"]