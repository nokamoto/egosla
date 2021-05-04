package service

import (
	"context"
	"log"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Watcher implements api.WatcherServiceServer.
type Watcher struct {
	api.UnimplementedWatcherServiceServer
	p persistent
	n nameGenerator
}

// NewWatcher creates a new Watcher.
func NewWatcher(p persistent) *Watcher {
	return &Watcher{
		n: watcherNameGenerator{},
		p: p,
	}
}

func (w *Watcher) CreateWatcher(ctx context.Context, req *api.CreateWatcherRequest) (*api.Watcher, error) {
	validate := func(req *api.CreateWatcherRequest) error {
		return nil
	}
	
	err := validate(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}

	created := req.GetWatcher()
	created.Name = w.n.newName()

	err = w.p.Create(created)
	if err != nil {
		log.Printf("unavailable: %s", err)
		return nil, status.Errorf(codes.Unavailable, "unavailable")
	}

	return created, nil
}
