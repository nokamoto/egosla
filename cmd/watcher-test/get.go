package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap"
)

func testGet(c api.WatcherServiceClient) test.Scenario {
	return test.Scenario{
		Name: "GetWatcher",
		Run: func(state test.State, logger *zap.Logger) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var created api.Watcher
			if err := state.Get(createdRecord, &created); err != nil {
				return err
			}

			res, err := c.GetWatcher(ctx, &api.GetWatcherRequest{
				Name: created.GetName(),
			})
			if err != nil {
				return err
			}

			if err := prototest.Equal(&created, res); err != nil {
				return err
			}
			return nil
		},
	}
}
