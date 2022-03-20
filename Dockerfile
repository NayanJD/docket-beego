FROM library/golang:1.17

ENV RUN_MODE=prod
ENV PORT=8080
ENV DB_CONN_STR=postgres://beego:password@localhost:5432/docket_local2?sslmode=disable
ENV APP_DIR $GOPATH/Users/nayan/Sites/projects/experiments/go/docket-beego

EXPOSE 8080

WORKDIR $APP_DIR

ADD . $APP_DIR

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

# Compile the binary and statically link
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s'

RUN chmod +x ./startup.sh

RUN go get github.com/beego/bee/v2

# Set the entrypoint
ENTRYPOINT [ "./startup.sh" ]




