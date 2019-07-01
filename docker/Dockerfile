FROM golang:latest

ADD . /go/src/github.com/Fedkasin/users-api

WORKDIR /go/src/github.com/Fedkasin/users-api

EXPOSE 8080

RUN go install github.com/Fedkasin/users-api

CMD /go/bin/users-api
