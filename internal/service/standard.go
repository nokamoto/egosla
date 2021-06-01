package service

import (
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createMethod(logger *zap.Logger, validate func() error, create func() error) error {
	logger.Debug("receive")

	err := validate()
	if err != nil {
		logger.Debug("invalid argument", zap.Error(err))
		return status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	err = create()
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return status.Errorf(codes.Unavailable, "unavailable")
	}

	return nil
}

type listRequest interface {
	GetPageToken() string
	GetPageSize() int32
}

func listMethod(logger *zap.Logger, req listRequest, list func(int, int) (int, error)) (string, error) {
	logger.Debug("receive")

	offset, err := fromPageToken(req.GetPageToken())
	if err != nil {
		logger.Debug("invalid argument", zap.Error(err))
		return "", status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	size := int(req.GetPageSize())
	if size <= 0 {
		size = defaultPageSize
	}

	got, err := list(offset, size+1)
	if err != nil {
		logger.Error("unavailable", zap.Error(err))
		return "", status.Errorf(codes.Unavailable, "unavailable")
	}

	var nextPageToken string
	if got == size+1 {
		nextPageToken = fromPageOffset(offset + size)
	}

	return nextPageToken, nil
}
