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
	p      persistent
	n      nameGenerator
	logger *zap.Logger
}

// NewWatcher creates a new Watcher.
func NewWatcher(p persistent, logger *zap.Logger) *Watcher {
	return &Watcher{
		n:      watcherNameGenerator{},
		p:      p,
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

	created := &api.Watcher{
		Name:     w.n.newName(),
		Keywords: req.GetWatcher().GetKeywords(),
	}

	err = w.p.Create(created)
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return nil, status.Errorf(codes.Unavailable, "unavailable")
	}

	logger.Info("created", zap.Any("watcher", created))

	return created, nil
}

func (w *Watcher) ListWatcher(ctx context.Context, req *api.ListWatcherRequest) (*api.ListWatcherResponse, error) {
	validate := func(req *api.ListWatcherRequest) error {
		_, err := fromPageToken(req.GetPageToken())
		return err
	}

	logger := w.logger.With(zap.Any("req", req), zap.String("method", "ListWatcher"))
	logger.Debug("receive")

	err := validate(req)
	if err != nil {
		logger.Debug("invalid argument", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	offset, _ := fromPageToken(req.GetPageToken())
	size := int(req.GetPageSize())
	if size <= 0 {
		size = defaultPageSize
	}

	watchers, err := w.p.List(offset, size+1)
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return nil, status.Errorf(codes.Unavailable, "unavailable")
	}

	var nextPageToken string
	if len(watchers) == size+1 {
		nextPageToken = fromPageOffset(offset + size)
		watchers = watchers[:len(watchers)-1]
	}

	return &api.ListWatcherResponse{
		NextPageToken: nextPageToken,
		Watchers:      watchers,
	}, nil
}

func (w *Watcher) DeleteWatcher(ctx context.Context, req *api.DeleteWatcherRequest) (*empty.Empty, error) {
	validate := func(req *api.DeleteWatcherRequest) error {
		return nil
	}

	logger := w.logger.With(zap.Any("req", req), zap.String("method", "DeleteWatcher"))
	logger.Debug("receive")

	err := validate(req)
	if err != nil {
		logger.Debug("invalid argument", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	err = w.p.Delete(req.GetName())
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return nil, status.Errorf(codes.Unavailable, "unavailable")
	}

	return &empty.Empty{}, nil
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
