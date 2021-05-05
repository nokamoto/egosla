package main

import (
	"log"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/os"
	"google.golang.org/grpc"
)

const (
	watcherAddress = "WATCHER_ADDRESS"
)

func main() {
	address := os.GetenvOr(watcherAddress, "127.0.0.1:9000")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewWatcherServiceClient(conn)

	scenarios{
		testCreate(c),
	}.run()
}
