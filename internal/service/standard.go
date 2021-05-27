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
