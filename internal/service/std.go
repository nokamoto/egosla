package service

import (
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
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

type getRequest interface {
	GetName() string
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

type listRequest interface {
	GetPageToken() string
	GetPageSize() int32
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

type deleteRequest interface {
	GetName() string
}

func (s *std) delete(validate func() error, req deleteRequest) (*empty.Empty, error) {
	l := s.logger.With(zap.String("method", "delete"), zap.String("name", req.GetName()))
	l.Debug("delete")

	if err := validate(); err != nil {
		l.Debug("invalid argument", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.persistent.Delete(req.GetName()); err != nil {
		l.Error("unknown", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "unavailable")
	}

	return &empty.Empty{}, nil
}
