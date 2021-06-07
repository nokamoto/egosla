package main

import (
	"context"
	"fmt"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd/test"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap"
)

func testList(c api.SubscriptionServiceClient) test.Scenario {
	return test.Scenario{
		Name: "ListSubscription",
		Run: func(state test.State, logger *zap.Logger) error {
			var expected api.Subscription
			if err := state.Get(createdRecord, &expected); err != nil {
				return err
			}

			return test.Until(func(ctx context.Context, pageToken string, pageSize int32) (test.ListResponse, bool, error) {
				res, err := c.ListSubscription(ctx, &api.ListSubscriptionRequest{
					PageToken: pageToken,
					PageSize:  pageSize,
				})
				if len(res.GetSubscriptions()) != 1 {
					return nil, false, fmt.Errorf("unexpected responnse: %v", res)
				}

				return res, prototest.Equal(&expected, res.GetSubscriptions()[0]) == nil, err
			})
		},
	}
}
