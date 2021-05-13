package main

import (
	"context"
	"errors"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func testUpdate(c api.WatcherServiceClient) scenario {
	return scenario{
		name: "UpdateWatcher",
		run: func(s state, logger *zap.Logger) (state, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var created api.Watcher
			err := prototext.Unmarshal([]byte(s[createdRecord]), &created)
			if err != nil {
				return nil, err
			}

			keywords := []string{"baz", "qux"}

			updateMask, err := fieldmaskpb.New(&created, "keywords")
			if err != nil {
				return nil, err
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
				return nil, err
			}

			logger.Info("updated", zap.Any("updated", updated))

			expected := &api.Watcher{
				Name:     created.GetName(),
				Keywords: keywords,
			}
			if diff := cmp.Diff(expected, updated, protocmp.Transform()); len(diff) != 0 {
				return nil, errors.New(diff)
			}

			s[createdRecord] = prototext.Format(updated)

			return s, nil
		},
	}
}
