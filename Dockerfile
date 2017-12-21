FROM golang:1.8
COPY . "G:/GoWorks/src/agenda-api-first"
RUN cd "$GOPATH/src/agenda-api-first/cli" && go get -v && go install -v
RUN cd "$GOPATH/src/agenda-api-first/service" && go get -v && go install -v
WORKDIR /
EXPOSE 8080
VOLUME ["/data"]
