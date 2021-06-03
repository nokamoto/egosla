package main

import (
	"context"
	"time"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"go.uber.org/zap"
)

func testDelete(c api.SubscriptionServiceClient) test.Scenario {
	return test.Scenario{
		Name: "DeleteSubscription",
		Run: func(state test.State, logger *zap.Logger) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var deleted api.Subscription
			if err := state.Get(createdRecord, &deleted); err != nil {
				return err
			}

			_, err := c.DeleteSubscription(ctx, &api.DeleteSubscriptionRequest{
				Name: deleted.GetName(),
			})
			if err != nil {
				return err
			}

			logger.Info("deleted", zap.String("name", deleted.GetName()))

			state.Delete(createdRecord)

			return nil
		},
	}
}
