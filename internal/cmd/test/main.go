package test

import (
	"github.com/nokamoto/egosla/internal/cmd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	// GrpcAddress is optional.
	GrpcAddress = "GRPC_ADDRESS"
)

// Test tests with a single gRPC client connection.
func Test(defaultAddresss string, f func(*grpc.ClientConn) Scenarios) {
	logger := cmd.NewLogger(true)
	defer logger.Sync()

	address := cmd.GetenvOr(GrpcAddress, defaultAddresss)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatal("did not connect", zap.Error(err), zap.String("address", address))
	}
	defer conn.Close()

	logger.Info("connect", zap.String("address", address))

	f(conn).Run(logger)
}
