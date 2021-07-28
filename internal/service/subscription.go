package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
)

// Subscription implements api.SubscriptionServiceServer.
type Subscription struct {
	api.UnimplementedSubscriptionServiceServer
	std    *std
	naming nameGenerator
}

// NewSubscription creates a new Subscription.
func NewSubscription(p persistent, logger *zap.Logger) *Subscription {
	return &Subscription{
		std: &std{
			logger:     logger,
			persistent: p,
		},
		naming: newSubscriptionNameGenerator(),
	}
}

func (s *Subscription) revert(v proto.Message) (*api.Subscription, error) {
	res, ok := v.(*api.Subscription)
	if !ok {
		s.std.logger.Error("[bug] unknown value", zap.Any("value", v))
		return nil, status.Error(codes.Internal, "internal error occurred")
	}
	return res, nil
}

func (s *Subscription) CreateSubscription(ctx context.Context, req *api.CreateSubscriptionRequest) (*api.Subscription, error) {
	created := &api.Subscription{
		Name:    s.naming.newName(),
		Watcher: req.GetSubscription().GetWatcher(),
	}
	if err := s.std.create(
		func() error {
			return nil
		},
		created,
	); err != nil {
		return nil, err
	}
	return created, nil
}

func (s *Subscription) GetSubscription(ctx context.Context, req *api.GetSubscriptionRequest) (*api.Subscription, error) {
	res, err := s.std.get(
		func() error {
			return nil
		},
		req,
	)
	if err != nil {
		return nil, err
	}
	return s.revert(res)
}

func (s *Subscription) ListSubscription(ctx context.Context, req *api.ListSubscriptionRequest) (*api.ListSubscriptionResponse, error) {
	xs, token, err := s.std.list(req)
	if err != nil {
		return nil, err
	}

	res := &api.ListSubscriptionResponse{
		NextPageToken: token,
	}
	for _, x := range xs {
		v, err := s.revert(x)
		if err != nil {
			return nil, err
		}
		res.Subscriptions = append(res.Subscriptions, v)
	}
	return res, nil
}

func (s *Subscription) UpdateSubscription(ctx context.Context, req *api.UpdateSubscriptionRequest) (*api.Subscription, error) {
	res, err := s.std.update(
		func() error {
			return nil
		},
		req,
		req.GetSubscription(),
	)
	if err != nil {
		return nil, err
	}
	return s.revert(res)
}

func (s *Subscription) DeleteSubscription(ctx context.Context, req *api.DeleteSubscriptionRequest) (*empty.Empty, error) {
	return s.std.delete(
		func() error {
			return nil
		},
		req,
	)
}
