package service

import (
	"context"
	"log"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type persistent interface {
	Create(*api.Watcher) error
}

// Watcher implements api.WatcherServiceServer.
type Watcher struct {
	api.UnimplementedWatcherServiceServer
	p persistent
}

func (w *Watcher) CreateWatcher(ctx context.Context, req *api.CreateWatcherRequest) (*api.Watcher, error) {
	validate := func(req *api.CreateWatcherRequest) error {
		return nil
	}
	
	err := validate(req)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	created := req.GetWatcher()

	err = w.p.Create(created)
	if err != nil {
		log.Printf("unavailable: %s", err)
		return nil, grpc.Errorf(codes.Unavailable, "unavailable")
	}

	return created, nil
}
