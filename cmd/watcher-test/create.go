package main

import (
	"context"
	"log"
	"time"

	"github.com/nokamoto/egosla/api"
)

func testCreate(c api.WatcherServiceClient) scenario {
	return scenario{
		name: "CreateWatcher",
		run: func(s state) (state, error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			res, err := c.CreateWatcher(ctx, &api.CreateWatcherRequest{
				Watcher: &api.Watcher{
					Keywords: []string{"foo", "bar"},
				},
			})
			if err != nil {
				return nil, err
			}

			log.Printf("got %v", res)

			return s, nil
		},
	}
}
