package main

import (
	"context"
	"fmt"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap"
)

func testList(c api.WatcherServiceClient) test.Scenario {
	return test.Scenario{
		Name: "ListWatcher",
		Run: func(s test.State, logger *zap.Logger) error {
			var expected api.Watcher
			if err := s.Get(createdRecord, &expected); err != nil {
				return err
			}

			return test.Until(func(ctx context.Context, pageToken string, pageSize int32) (test.ListResponse, bool, error) {
				res, err := c.ListWatcher(ctx, &api.ListWatcherRequest{
					PageToken: pageToken,
					PageSize:  pageSize,
				})

				if len(res.GetWatchers()) != 1 {
					return nil, false, fmt.Errorf("unexpected response: %v", res)
				}

				return res, prototest.Equal(&expected, res.GetWatchers()[0]) == nil, err
			})
		},
	}
}
