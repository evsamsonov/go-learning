FROM golang:1.15.1

WORKDIR /go/src/app
COPY ./hello_server.go hello_server.go

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
