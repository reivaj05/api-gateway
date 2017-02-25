FROM golang:alpine


WORKDIR /go/src/github.com/reivaj05/apigateway

ADD . /go/src/github.com/reivaj05/apigateway

RUN apk -U add make git bash wget gcc g++ protobuf
RUN make
RUN apk del make git curl gcc g++

ENTRYPOINT apigateway
