package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	testWatcherRecord = "test-watcher"
)

func withWatcher(f func(api.WatcherServiceClient) error) error {
	address := cmd.GetenvOr(watcherGrpcAddress, "127.0.0.1:9000")

	cc, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer cc.Close()

	return f(api.NewWatcherServiceClient(cc))
}

var setup = test.Scenario{
	Name: "setup: CreateWatcher",
	Run: func(state test.State, logger *zap.Logger) error {
		err := withWatcher(func(c api.WatcherServiceClient) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			res, err := c.CreateWatcher(ctx, &api.CreateWatcherRequest{})
			if err != nil {
				return err
			}

			state.Set(testWatcherRecord, res)

			return nil
		})
		return err
	},
}

var teardown = test.Scenario{
	Name: "teardown: DeleteWatcher",
	Run: func(state test.State, logger *zap.Logger) error {
		err := withWatcher(func(c api.WatcherServiceClient) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var deleted api.Watcher
			if err := state.Get(testWatcherRecord, &deleted); err != nil {
				return err
			}

			_, err := c.DeleteWatcher(ctx, &api.DeleteWatcherRequest{
				Name: deleted.GetName(),
			})
			if err != nil {
				return err
			}

			state.Delete(testWatcherRecord)

			return nil
		})
		return err
	},
}
