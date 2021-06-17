all: protoc patch format go yarn lint

patch:
	go get github.com/uber-go/gopatch
	gopatch -p ./scripts/gomock-v1.5.0.patch ./...

format:
	clang-format --style=Google -i api/*.proto
	go fmt ./cmd/...
	go fmt ./internal/...
	cd web && yarn && yarn format

lint:
	go get honnef.co/go/tools/cmd/staticcheck
	staticcheck ./... 
	cd web && yarn lint

protoc:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	api/*.proto
	go mod tidy

	rm -rf web/src/api/*
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

watcher:
	go get github.com/cespare/reflex
	GRPC_PORT=9000 reflex -r '(\.go|go\.mod|go\.sum)' -s go run ./cmd/watcher

subscription:
	go get github.com/cespare/reflex
	GRPC_PORT=9001 reflex -r '(\.go|go\.mod|go\.sum)' -s go run ./cmd/subscription
