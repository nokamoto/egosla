package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
)

func testList(c api.WatcherServiceClient) test.Scenario {
	return test.Scenario{
		Name: "ListWatcher",
		Run: func(s test.State, logger *zap.Logger) (test.State, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var expected api.Watcher
			if err := s.Get(createdRecord, &expected); err != nil {
				return nil, err
			}

			logger.Info("until", zap.Any("expected", &expected))

			var nextPageToken string
			for {
				logger.Info("request", zap.String("token", nextPageToken))
				res, err := c.ListWatcher(ctx, &api.ListWatcherRequest{
					PageToken: nextPageToken,
					PageSize:  1,
				})
				if err != nil {
					return nil, err
				}
				if len(res.GetWatchers()) != 1 {
					return nil, fmt.Errorf("unexpected response: %v", res)
				}

				if test.Equal(&expected, res.GetWatchers()[0]) == nil {
					break
				}

				logger.Info("request next", zap.Any("got", res))

				nextPageToken = res.GetNextPageToken()
			}

			logger.Info("found")

			return s, nil
		},
	}
}
