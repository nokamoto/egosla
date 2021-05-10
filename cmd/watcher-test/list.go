package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/protobuf/testing/protocmp"
)

func testList(c api.WatcherServiceClient) scenario {
	return scenario{
		name: "ListWatcher",
		run: func(s state, logger *zap.Logger) (state, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var expected api.Watcher
			err := proto.UnmarshalText(s[createdRecord], &expected)
			if err != nil {
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

				if diff := cmp.Diff(&expected, res.GetWatchers()[0], protocmp.Transform()); len(diff) == 0 {
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
