package main

import (
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"google.golang.org/grpc"
)

func main() {
	test.Test("127.0.0.1:9001", func(cc *grpc.ClientConn) test.Scenarios {
		_ = api.NewSubscriptionServiceClient(cc)
		return test.Scenarios{}
	})
}
