all: protoc format go

format:
	clang-format --style=Google -i api/*.proto
	go fmt ./cmd/...

protoc:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	api/*.proto

go:
	go test ./...
	go mod tidy