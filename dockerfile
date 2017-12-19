FROM golang:latest



WORKDIR $GOPATH/src/github.com/renleimlj/agenda-api-first/cli

ADD . $GOPATH/src/github.com/renleimlj/agenda-api-first

EXPOSE 8080

CMD  $GOPATH/src/github.com/renleimlj/agenda-api-first/main
