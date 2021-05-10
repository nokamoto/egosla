package main

import (
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	watcherAddress = "WATCHER_ADDRESS"
)

func main() {
	logger := cmd.NewLogger(true)
	defer logger.Sync()

	address := cmd.GetenvOr(watcherAddress, "127.0.0.1:9000")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatal("did not connect", zap.Error(err), zap.String("address", address))
	}
	defer conn.Close()

	logger.Info("connect", zap.String("address", address))

	c := api.NewWatcherServiceClient(conn)

	scenarios{
		testCreate(c),
		testList(c),
	}.run(logger)
}
