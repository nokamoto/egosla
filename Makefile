all: proto format

format:
	clang-format --style=Google -i api/*

proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc