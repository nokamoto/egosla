package service

import (
	"context"
	"errors"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
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
	l := s.logger.With(zap.String("method", "create"), zap.Any("resource", created))
	l.Debug("create")

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

func (s *std) get(validate func() error, req getRequest) (proto.Message, error) {
	l := s.logger.With(zap.String("method", "get"), zap.Any("name", req.GetName()))
	l.Debug("get")

	if err := validate(); err != nil {
		l.Debug("invalid argument", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := s.persistent.Get(req.GetName())
	if errors.Is(err, mysql.ErrNotFound) {
		l.Debug("not found", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "%s not found", req.GetName())
	}
	if err != nil {
		l.Error("unknown", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "unavailable")
	}

	return res, nil
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

func (w *WatcherV2) GetWatcher(ctx context.Context, req *api.GetWatcherRequest) (*api.Watcher, error) {
	res, err := w.std.get(func() error {
		return nil
	}, req)
	if err != nil {
		return nil, err
	}

	v, ok := res.(*api.Watcher)
	if !ok {
		w.std.logger.Error("[bug] unknown value", zap.Any("value", v))
		return nil, status.Error(codes.Internal, "internal error occurred")
	}

	return v, nil
}
