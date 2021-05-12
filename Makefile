all: protoc format go yarn

format:
	clang-format --style=Google -i api/*.proto
	go fmt ./cmd/...
	go fmt ./internal/...
	go mod tidy
	cd web && yarn && yarn format

protoc:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	api/*.proto
	go mod tidy

	protoc \
		--js_out=import_style=commonjs:web/src \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:web/src \
		api/*.proto

go:
	go get github.com/golang/mock/mockgen
	go generate ./...
	go test ./...
	go mod tidy

yarn:
	cd web && yarn test --watchAll=false
	cd web && yarn lint

watcher:
	go get github.com/cespare/reflex
	reflex -r '(\.go|go\.mod|go\.sum)' -s go run ./cmd/watcher
