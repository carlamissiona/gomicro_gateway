
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/bookgateway.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/bookgateway.proto
	
.PHONY: build
build:
	go build -o bookgateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t bookgateway:latest
