#!/bin/bash

set -x -e;

SERVICES_PATH=./protos/services
API_ENDPOINTS_PATH=./protos/api

# Usage spliteroo list separator return-index
# ex. spliteroo a.a b.a c.a . 1
#   -> a b c
# ex. spliteroo a.a b.a c.a . 2
#   -> a a a
spliteroo() {
	echo `echo $1| cut -d $2 -f $3`
}

generate_services() {
	protoc -I$SERVICES_PATH \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:"$GOPATH/src" \
		"$SERVICES_PATH/$1.proto"
}

generate_api_endpoints() {
	mkdir -p "./api/$1/generated/"
	protoc -I$API_ENDPOINTS_PATH \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:"$GOPATH/src" \
		"$API_ENDPOINTS_PATH/$1.proto"
	protoc -I$API_ENDPOINTS_PATH \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:"./api/$1/generated/" \
		"$API_ENDPOINTS_PATH/$1.proto"
	protoc -I$API_ENDPOINTS_PATH \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:"./api/swagger/" \
		"$API_ENDPOINTS_PATH/$1.proto"
}

for x in $(ls -1 "$SERVICES_PATH")
do
	generate_services $(spliteroo $x '.' 1) $OUT_PATH
done

mkdir -p api/swagger
for x in $(ls -1 "$API_ENDPOINTS_PATH" | grep -v 'google')
do
	generate_api_endpoints $(spliteroo $x '.' 1) $OUT_PATH
done
