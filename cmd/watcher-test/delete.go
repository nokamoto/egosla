package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
)

func testDelete(c api.WatcherServiceClient) test.Scenario {
	return test.Scenario{
		Name: "DeleteWatcher",
		Run: func(s test.State, logger *zap.Logger) (test.State, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var deleted api.Watcher
			if err := s.Get(createdRecord, &deleted); err != nil {
				return nil, err
			}

			_, err := c.DeleteWatcher(ctx, &api.DeleteWatcherRequest{
				Name: deleted.GetName(),
			})
			if err != nil {
				return nil, err
			}

			logger.Info("deleted", zap.String("name", deleted.GetName()))

			s.Delete(createdRecord)

			return s, nil
		},
	}
}
