GOCMD = go
PKG   = ./...

.PHONY: %

default: fmt deps proto test build

all: build
build: deps
	$(GOCMD) install
fmt:
	$(GOCMD) fmt $(PKG)
test: deps
	$(GOCMD) test -a -v ./...
deps:
	wget https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm
	chmod +x gpm
	./gpm
	rm gpm
protobuf:
	$(GOCMD) get -u github.com/golang/protobuf/proto
	$(GOCMD) get -u github.com/golang/protobuf/protoc-gen-go
	$(GOCMD) get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOCMD) get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
proto: protobuf
	./proto-gen.sh
