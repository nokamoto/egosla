package service

import (
	"context"

	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Watcher implements api.WatcherServiceServer.
type Watcher struct {
	api.UnimplementedWatcherServiceServer
	p persistent
	n nameGenerator
	logger *zap.Logger
}

// NewWatcher creates a new Watcher.
func NewWatcher(p persistent, logger *zap.Logger) *Watcher {
	return &Watcher{
		n: watcherNameGenerator{},
		p: p,
		logger: logger.With(zap.String("service", "WatcherService")),
	}
}

func (w *Watcher) CreateWatcher(ctx context.Context, req *api.CreateWatcherRequest) (*api.Watcher, error) {
	validate := func(req *api.CreateWatcherRequest) error {
		return nil
	}

	logger := w.logger.With(zap.Any("req", req), zap.String("method", "CreateWatcher"))
	logger.Debug("receive")
	
	err := validate(req)
	if err != nil {
		logger.Debug("invalid argument", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	created := req.GetWatcher()
	created.Name = w.n.newName()

	err = w.p.Create(created)
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return nil, status.Errorf(codes.Unavailable, "unavailable")
	}

	return created, nil
}
