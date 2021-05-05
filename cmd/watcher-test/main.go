package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/grpc"
)

const (
	watcherAddress = "WATCHER_ADDRESS"
)

func main() {
	conn, err := grpc.Dial(os.Getenv(watcherAddress), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewWatcherServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.CreateWatcher(ctx, &api.CreateWatcherRequest{
		Watcher: &api.Watcher{
			Keywords: []string{"foo", "bar"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got %v", res)
}
