package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/cmd"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlUser = "MYSQL_USER"
	mysqlPass = "MYSQL_PASS"
	mysqlTCP  = "MYSQL_TCP"
	mysqlDB   = "MYSQL_DB"
)

func main() {
	logger := cmd.NewLogger(cmd.GetenvDebug())
	defer logger.Sync()

	port := ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	secret := cmd.GetenvOr(mysqlPass, "")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		cmd.GetenvOr(mysqlUser, "root"),
		secret,
		cmd.GetenvOr(mysqlTCP, "127.0.0.1:3306"),
		cmd.GetenvOr(mysqlDB, "egosla"),
	)
	printableDsn := strings.ReplaceAll(dsn, secret, strings.Repeat("*", len(secret)))

	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to mysql.Open", zap.Error(err), zap.String("dsn", printableDsn))
	}

	logger.Info("mysql.Open", zap.String("dsn", printableDsn))

	s := grpc.NewServer()
	api.RegisterWatcherServiceServer(s, service.NewWatcher(mysql.NewPersistentWatcher(db), logger))
	if err := s.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
