package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/service"
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
	port := ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv(mysqlUser),
		os.Getenv(mysqlPass),
		os.Getenv(mysqlTCP),
		os.Getenv(mysqlDB),
	)
	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to mysql.Open: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterWatcherServiceServer(s, service.NewWatcher(mysql.NewPersistentWatcher(db)))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
