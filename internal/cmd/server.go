package cmd

import (
	"fmt"
	"net"
	"os"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// GrpcPort is required. It should be a port number (e.g. 9000).
	GrpcPort = "GRPC_PORT"
)

const (
	// MysqlUser is optional (default is root).
	MysqlUser = "MYSQL_USER"
	// MysqlPass is optional (default is "").
	MysqlPass = "MYSQL_PASS"
	// MysqlTCP is optional (default is 127.0.0.1:3306).
	MysqlTCP = "MYSQL_TCP"
	// MysqlDB is optional (default is egosla).
	MysqlDB = "MYSQL_DB"
)

// GrpcServer serves a new gRPC server.
// It calls `register` to register gRPC services with a mysql connection and a logger.
func GrpcServer(register func(*grpc.Server, *gorm.DB, *zap.Logger) error) {
	logger := NewLogger(GetenvDebug())
	defer logger.Sync()

	port := fmt.Sprintf(":%s", os.Getenv(GrpcPort))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err), zap.String("port", port))
	}

	logger.Info("listen", zap.String("port", port), zap.String("address", lis.Addr().String()))

	secret := GetenvOr(MysqlPass, "")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		GetenvOr(MysqlUser, "root"),
		secret,
		GetenvOr(MysqlTCP, "127.0.0.1:3306"),
		GetenvOr(MysqlDB, "egosla"),
	)
	printableDsn := strings.ReplaceAll(dsn, secret, strings.Repeat("*", len(secret)))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to mysql.Open", zap.Error(err), zap.String("dsn", printableDsn))
	}

	logger.Info("mysql.Open", zap.String("dsn", printableDsn))

	s := grpc.NewServer()

	if err := register(s, db, logger); err != nil {
		logger.Fatal("failed to register service(s)", zap.Error(err))
	}

	if err := s.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
