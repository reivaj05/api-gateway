FROM golang:alpine


WORKDIR /go/src/github.com/reivaj05/apigateway

ADD . /go/src/github.com/reivaj05/apigateway

RUN apk -U add make git bash wget curl gcc g++ protobuf protobuf-dev
RUN make
RUN apk del make git wget curl gcc g++ protobuf protobuf-dev

ENTRYPOINT apigateway
