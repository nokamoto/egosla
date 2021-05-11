package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/prototext"
)

const createdRecord = "created-record"

func testCreate(c api.WatcherServiceClient) scenario {
	return scenario{
		name: "CreateWatcher",
		run: func(s state, logger *zap.Logger) (state, error) {
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

			s[createdRecord] = prototext.Format(res)

			return s, nil
		},
	}
}
