package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
)

// Subscription implements api.SubscriptionServiceServer.
type Subscription struct {
	api.UnimplementedSubscriptionServiceServer
	p      persistentSubscription
	n      nameGenerator
	logger *zap.Logger
}

// NewSubscription creates a new Subscription.
func NewSubscription(p persistentSubscription, logger *zap.Logger) *Subscription {
	return &Subscription{
		n:      newSubscriptionNameGenerator(),
		p:      p,
		logger: logger.With(zap.String("service", "SubscriptionService")),
	}
}

func (s *Subscription) CreateSubscription(ctx context.Context, req *api.CreateSubscriptionRequest) (*api.Subscription, error) {
	created := &api.Subscription{
		Name:    s.n.newName(),
		Watcher: req.GetSubscription().GetWatcher(),
	}

	err := createMethod(
		s.logger.With(zap.Any("req", req), zap.String("method", "CreateSubscription"), zap.Any("created", created)),
		func() error {
			return nil
		},
		func() error {
			return s.p.Create(created)
		},
	)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Subscription) ListSubscription(ctx context.Context, req *api.ListSubscriptionRequest) (*api.ListSubscriptionResponse, error) {
	var res api.ListSubscriptionResponse
	nextPageToken, err := listMethod(
		s.logger.With(zap.Any("req", req), zap.String("method", "ListSubscription")),
		req,
		func(offset, limit int) (int, error) {
			v, err := s.p.List(offset, limit)
			res.Subscriptions = v
			return len(v), err
		},
	)
	if err != nil {
		return nil, err
	}

	res.NextPageToken = nextPageToken
	if len(nextPageToken) != 0 {
		res.Subscriptions = res.Subscriptions[:len(res.Subscriptions)-1]
	}

	return &res, nil
}

func (s *Subscription) DeleteSubscription(ctx context.Context, req *api.DeleteSubscriptionRequest) (*empty.Empty, error) {
	return deleteMethod(
		s.logger.With(zap.Any("req", req), zap.String("method", "DeleteSubscription")),
		s.p,
		req,
		func(_ string) error {
			return nil
		},
	)
}
