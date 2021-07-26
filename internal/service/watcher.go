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

// Watcher implements api.WatcherServiceServer.
type Watcher struct {
	api.UnimplementedWatcherServiceServer
	std    *std
	naming nameGenerator
}

// NewWatcher creates a new Watcher.
func NewWatcher(p persistent, logger *zap.Logger) *Watcher {
	return &Watcher{
		std: &std{
			logger:     logger.With(zap.String("service", "WatcherService")),
			persistent: p,
		},
		naming: newWatcherNameGenerator(),
	}
}

func (w *Watcher) revert(v proto.Message) (*api.Watcher, error) {
	res, ok := v.(*api.Watcher)
	if !ok {
		w.std.logger.Error("[bug] unknown value", zap.Any("value", v))
		return nil, status.Error(codes.Internal, "internal error occurred")
	}
	return res, nil
}

func (w *Watcher) CreateWatcher(ctx context.Context, req *api.CreateWatcherRequest) (*api.Watcher, error) {
	created := &api.Watcher{
		Name:     w.naming.newName(),
		Keywords: req.GetWatcher().GetKeywords(),
	}
	if err := w.std.create(func() error {
		return nil
	}, created); err != nil {
		return nil, err
	}
	return created, nil
}

func (w *Watcher) GetWatcher(ctx context.Context, req *api.GetWatcherRequest) (*api.Watcher, error) {
	res, err := w.std.get(func() error {
		return nil
	}, req)
	if err != nil {
		return nil, err
	}

	return w.revert(res)
}

func (w *Watcher) ListWatcher(ctx context.Context, req *api.ListWatcherRequest) (*api.ListWatcherResponse, error) {
	xs, token, err := w.std.list(req)
	if err != nil {
		return nil, err
	}

	res := &api.ListWatcherResponse{
		NextPageToken: token,
	}
	for _, x := range xs {
		v, err := w.revert(x)
		if err != nil {
			return nil, err
		}
		res.Watchers = append(res.Watchers, v)
	}
	return res, nil
}

func (w *Watcher) UpdateWatcher(ctx context.Context, req *api.UpdateWatcherRequest) (*api.Watcher, error) {
	res, err := w.std.update(
		func() error {
			return nil
		},
		req,
		req.GetWatcher(),
	)
	if err != nil {
		return nil, err
	}
	return w.revert(res)
}

func (w *Watcher) DeleteWatcher(ctx context.Context, req *api.DeleteWatcherRequest) (*empty.Empty, error) {
	return w.std.delete(
		func() error {
			return nil
		},
		req,
	)
}
