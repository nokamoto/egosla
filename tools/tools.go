// +build tools

package tools

import (
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/cespare/reflex"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "github.com/uber-go/gopatch"
)
