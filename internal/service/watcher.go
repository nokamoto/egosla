package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Watcher implements api.WatcherServiceServer.
type Watcher struct {
	api.UnimplementedWatcherServiceServer
	p      persistentWatcher
	n      nameGenerator
	logger *zap.Logger
}

// NewWatcher creates a new Watcher.
func NewWatcher(p persistentWatcher, logger *zap.Logger) *Watcher {
	return &Watcher{
		n:      newWatcherNameGenerator(),
		p:      p,
		logger: logger.With(zap.String("service", "WatcherService")),
	}
}

func (w *Watcher) CreateWatcher(ctx context.Context, req *api.CreateWatcherRequest) (*api.Watcher, error) {
	created := &api.Watcher{
		Name:     w.n.newName(),
		Keywords: req.GetWatcher().GetKeywords(),
	}

	err := createMethod(
		w.logger.With(zap.Any("req", req), zap.String("method", "CreateWatcher"), zap.Any("created", created)),
		func() error {
			return nil
		},
		func() error {
			return w.p.Create(created)
		},
	)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (w *Watcher) ListWatcher(ctx context.Context, req *api.ListWatcherRequest) (*api.ListWatcherResponse, error) {
	var res api.ListWatcherResponse

	nextPageToken, err := listMethod(
		w.logger.With(zap.Any("req", req), zap.String("method", "ListWatcher")),
		req,
		func(offset, limit int) (int, error) {
			v, err := w.p.List(offset, limit)
			res.Watchers = v
			return len(v), err
		},
	)
	if err != nil {
		return nil, err
	}

	res.NextPageToken = nextPageToken
	if len(nextPageToken) != 0 {
		res.Watchers = res.Watchers[:len(res.Watchers)-1]
	}

	return &res, nil
}

func (w *Watcher) DeleteWatcher(ctx context.Context, req *api.DeleteWatcherRequest) (*empty.Empty, error) {
	return deleteMethod(
		w.logger.With(zap.Any("req", req), zap.String("method", "DeleteWatcher")),
		w.p,
		req,
		func(_ string) error {
			return nil
		},
	)
}

func (w *Watcher) UpdateWatcher(ctx context.Context, req *api.UpdateWatcherRequest) (*api.Watcher, error) {
	validate := func(req *api.UpdateWatcherRequest) error {
		if !req.GetUpdateMask().IsValid(req.GetWatcher()) {
			return fmt.Errorf("invalid update mask: %v", req.GetUpdateMask())
		}
		for _, p := range req.GetUpdateMask().GetPaths() {
			if p == "name" {
				return errors.New("update name unsupported")
			}
		}
		return nil
	}

	logger := w.logger.With(zap.Any("req", req), zap.String("method", "UpdateWatcher"))
	logger.Debug("receive")

	err := validate(req)
	if err != nil {
		logger.Debug("invalid argument", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	set := &api.Watcher{
		Name:     req.GetName(),
		Keywords: req.GetWatcher().GetKeywords(),
	}

	updated, err := w.p.Update(set, req.GetUpdateMask())
	if errors.Is(err, mysql.ErrInvalidArgument) {
		logger.Debug("invalid argument", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return nil, status.Errorf(codes.Unavailable, "unavailable")
	}

	return updated, nil
}
