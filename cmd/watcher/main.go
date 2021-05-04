package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/service"
	"google.golang.org/grpc"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	port := ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname")
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
