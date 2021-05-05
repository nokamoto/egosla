package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

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

func getenvOr(key string, or string) string {
	s := os.Getenv(key)
	if len(s) == 0 {
		return or
	}
	return s
}

func main() {
	port := ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	secret := getenvOr(mysqlPass, "")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		getenvOr(mysqlUser, "root"),
		secret,
		getenvOr(mysqlTCP, "127.0.0.1:3306"),
		getenvOr(mysqlDB, "egosla"),
	)
	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to mysql.Open: %v", err)
	}

	log.Printf("mysql.Open: %v", strings.ReplaceAll(dsn, secret, strings.Repeat("*", len(secret))))

	s := grpc.NewServer()
	api.RegisterWatcherServiceServer(s, service.NewWatcher(mysql.NewPersistentWatcher(db)))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
