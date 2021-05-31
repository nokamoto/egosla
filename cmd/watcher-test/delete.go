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
		Run: func(s test.State, logger *zap.Logger) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var deleted api.Watcher
			if err := s.Get(createdRecord, &deleted); err != nil {
				return err
			}

			_, err := c.DeleteWatcher(ctx, &api.DeleteWatcherRequest{
				Name: deleted.GetName(),
			})
			if err != nil {
				return err
			}

			logger.Info("deleted", zap.String("name", deleted.GetName()))

			s.Delete(createdRecord)

			return nil
		},
	}
}
