package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/prototext"
)

func testDelete(c api.WatcherServiceClient) scenario {
	return scenario{
		name: "DeleteWatcher",
		run: func(s state, logger *zap.Logger) (state, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var deleted api.Watcher
			err := prototext.Unmarshal([]byte(s[createdRecord]), &deleted)
			if err != nil {
				return nil, err
			}

			_, err = c.DeleteWatcher(ctx, &api.DeleteWatcherRequest{
				Name: deleted.GetName(),
			})
			if err != nil {
				return nil, err
			}

			logger.Info("deleted", zap.String("name", deleted.GetName()))

			delete(s, createdRecord)

			return s, nil
		},
	}
}
