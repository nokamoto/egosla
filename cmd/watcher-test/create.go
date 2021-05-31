package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
)

const createdRecord = "created-record"

func testCreate(c api.WatcherServiceClient) test.Scenario {
	return test.Scenario{
		Name: "CreateWatcher",
		Run: func(s test.State, logger *zap.Logger) (test.State, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			res, err := c.CreateWatcher(ctx, &api.CreateWatcherRequest{
				Watcher: &api.Watcher{
					Keywords: []string{"foo", "bar"},
				},
			})
			if err != nil {
				return nil, err
			}

			logger.Info("got", zap.Any("res", res))

			s.Set(createdRecord, res)

			return s, nil
		},
	}
}
