package test

import (
	"context"
	"fmt"
	"time"
)

type ListResponse interface {
	GetNextPageToken() string
}

// Until iterates the list method until it returns an expected resource.
func Until(f func(context.Context, string, int32) (ListResponse, bool, error)) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var nextPageToken string
	for {
		res, ok, err := f(ctx, nextPageToken, 1)
		if err != nil {
			return fmt.Errorf("%s: %w", nextPageToken, err)
		}
		if ok {
			return nil
		}

		nextPageToken = res.GetNextPageToken()
		if len(nextPageToken) == 0 {
			break
		}
	}

	return fmt.Errorf("unmet expectation")
}
