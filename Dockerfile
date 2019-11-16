FROM golang:latest

ADD . /go/src/github.com/chinsyo/quarxlab

RUN go mod tidy && go mod vendor

ENTRYPOINT /go/bin/quarxlab

EXPOSE 8000
