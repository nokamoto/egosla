package main

import (
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"google.golang.org/grpc"
)

const (
	watcherGrpcAddress = "WATCHER_GRPC_ADDRESS"
)

func main() {
	test.Test(
		"127.0.0.1:9001",
		func(cc *grpc.ClientConn) test.Scenarios {
			c := api.NewSubscriptionServiceClient(cc)
			return test.Scenarios{
				testCreate(c),
			}
		},
		&setup,
		&teardown,
	)
}
