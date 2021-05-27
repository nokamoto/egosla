package service

import (
	"context"

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

// NewWatcher creates a new Watcher.
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
