FROM golang:alpine

COPY . /go/src/github.com/lordralex/mightyena

RUN \
    go get github.com/lordralex/mightyena
    go install github.com/lordralex/mightyena && \
    rm -rf /go/src && \
    apk add --no-cache py-pip && \
    pip install mcstatus

CMD ["mightyena"]
