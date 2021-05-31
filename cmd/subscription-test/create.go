package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
)

const (
	createdRecord = "created-subscription"
)

func testCreate(c api.SubscriptionServiceClient) test.Scenario {
	return test.Scenario{
		Name: "CreateSubscription",
		Run: func(state test.State, logger *zap.Logger) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var watcher api.Watcher
			if err := state.Get(testWatcherRecord, &watcher); err != nil {
				return err
			}

			res, err := c.CreateSubscription(ctx, &api.CreateSubscriptionRequest{
				Subscription: &api.Subscription{
					Watcher: watcher.GetName(),
				},
			})
			if err != nil {
				return err
			}

			logger.Info("got", zap.Any("res", res))

			state.Set(createdRecord, res)

			return nil
		},
	}
}
