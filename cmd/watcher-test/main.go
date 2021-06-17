package main

import (
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"google.golang.org/grpc"
)

func main() {
	test.Test(
		"127.0.0.1:9000",
		func(cc *grpc.ClientConn) test.Scenarios {
			c := api.NewWatcherServiceClient(cc)
			return test.Scenarios{
				testCreate(c),
				testGet(c),
				testList(c),
				testUpdate(c),
				testDelete(c),
			}
		},
		nil,
		nil,
	)
}
