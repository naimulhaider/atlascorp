FROM golang

ADD . /go/src/github.com/naimulhaider/atlascorp

RUN go install github.com/naimulhaider/atlascorp/cmd/dns

ENTRYPOINT ["/go/bin/dns"]