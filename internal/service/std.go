package service

import (
	"context"
	"errors"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"go.uber.org/zap"
	"google.golang.org/genproto/protobuf/field_mask"
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
	l := s.logger.With(zap.String("method", "get"), zap.String("name", req.GetName()))
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

func (s *std) list(req listRequest) ([]proto.Message, string, error) {
	l := s.logger.With(zap.String("method", "list"), zap.String("token", req.GetPageToken()), zap.Int32("size", req.GetPageSize()))
	l.Debug("list")

	offset, err := fromPageToken(req.GetPageToken())
	if err != nil {
		l.Debug("invalid argument", zap.Error(err))
		return nil, "", status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	size := int(req.GetPageSize())
	if size <= 0 {
		size = defaultPageSize
	}

	got, err := s.persistent.List(offset, size+1)
	if err != nil {
		l.Error("unavailable", zap.Error(err))
		return nil, "", status.Errorf(codes.Unavailable, "unavailable")
	}

	var nextPageToken string
	if len(got) == size+1 {
		nextPageToken = fromPageOffset(offset + size)
		got = got[:len(got)-1]
	}

	return got, nextPageToken, nil
}

type updateRequest interface {
	GetName() string
	GetUpdateMask() *field_mask.FieldMask
}

func (s *std) update(validate func() error, req updateRequest, value proto.Message) (proto.Message, error) {
	l := s.logger.With(zap.String("method", "update"), zap.String("name", req.GetName()))
	l.Debug("update")

	if err := validate(); err != nil {
		l.Debug("invalid argument", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := s.persistent.Update(req.GetName(), req.GetUpdateMask(), value)
	if errors.Is(err, mysql.ErrInvalidArgument) {
		l.Debug("invalid argument", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
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

func (w *WatcherV2) revert(v proto.Message) (*api.Watcher, error) {
	res, ok := v.(*api.Watcher)
	if !ok {
		w.std.logger.Error("[bug] unknown value", zap.Any("value", v))
		return nil, status.Error(codes.Internal, "internal error occurred")
	}
	return res, nil
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

	return w.revert(res)
}

func (w *WatcherV2) ListWatcher(ctx context.Context, req *api.ListWatcherRequest) (*api.ListWatcherResponse, error) {
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

func (w *WatcherV2) UpdateWatcher(ctx context.Context, req *api.UpdateWatcherRequest) (*api.Watcher, error) {
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
