package main

import (
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cmd.GrpcServer(func(s *grpc.Server, db *gorm.DB, logger *zap.Logger) error {
		api.RegisterSubscriptionServiceServer(
			s,
			service.NewSubscription(
				mysql.NewStdPersistent(db, &mysql.SubscriptionModel{}),
				logger,
			),
		)
		return nil
	})
}
