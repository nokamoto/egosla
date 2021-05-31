package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func testUpdate(c api.WatcherServiceClient) test.Scenario {
	return test.Scenario{
		Name: "UpdateWatcher",
		Run: func(s test.State, logger *zap.Logger) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var created api.Watcher
			if err := s.Get(createdRecord, &created); err != nil {
				return err
			}

			keywords := []string{"baz", "qux"}

			updateMask, err := fieldmaskpb.New(&created, "keywords")
			if err != nil {
				return err
			}

			logger.Info("update", zap.Strings("keywords", keywords), zap.Any("updateMask", updateMask))

			updated, err := c.UpdateWatcher(ctx, &api.UpdateWatcherRequest{
				Name: created.GetName(),
				Watcher: &api.Watcher{
					Keywords: keywords,
				},
				UpdateMask: updateMask,
			})
			if err != nil {
				return err
			}

			logger.Info("updated", zap.Any("updated", updated))

			expected := &api.Watcher{
				Name:     created.GetName(),
				Keywords: keywords,
			}
			if err := test.Equal(expected, updated); err != nil {
				return err
			}

			s.Set(createdRecord, updated)

			return nil
		},
	}
}
