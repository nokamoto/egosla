package service

import (
	"context"

	"github.com/nokamoto/egosla/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type std struct {
	logger     *zap.Logger
	persistent persistent
}

func (s *std) create(validate func() error, created proto.Message) error {
	l := s.logger.With(zap.String("method", "create"))
	l.Debug("create", zap.Any("resource", created))

	if err := validate(); err != nil {
		l.Debug("invalid argument", zap.Error(err))
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.persistent.Create(created); err != nil {
		l.Error("unknown", zap.Error(err))
		return status.Error(codes.Unavailable, "unavailable")
	}

	return nil
}

type WatcherV2 struct {
	api.UnimplementedWatcherServiceServer
	std    *std
	naming nameGenerator
}

func (w *WatcherV2) CreateWatcher(ctx context.Context, req *api.CreateWatcherRequest) (*api.Watcher, error) {
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
