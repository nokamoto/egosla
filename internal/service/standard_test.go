package service

import (
	"context"
	"testing"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestWatcher_ListWatcher_listMethod(t *testing.T) {
	elm1 := &api.Watcher{
		Name: "foo",
	}

	elm2 := &api.Watcher{
		Name: "bar",
	}

	testcases := []struct {
		name     string
		req      *api.ListWatcherRequest
		mock     func(p *MockpersistentWatcher, _ *MocknameGenerator)
		expected *api.ListWatcherResponse
		code     codes.Code
	}{
		{
			name: "ok: len(resources) < page size",
			req: &api.ListWatcherRequest{
				PageToken: "10",
				PageSize:  1,
			},
			mock: func(p *MockpersistentWatcher, _ *MocknameGenerator) {
				p.EXPECT().List(10, 2).Return(nil, nil)
			},
			expected: &api.ListWatcherResponse{
				Watchers: []*api.Watcher{},
			},
		},
		{
			name: "ok: len(resources) == page size",
			req: &api.ListWatcherRequest{
				PageSize: 1,
			},
			mock: func(p *MockpersistentWatcher, _ *MocknameGenerator) {
				p.EXPECT().List(0, 2).Return([]*api.Watcher{elm1}, nil)
			},
			expected: &api.ListWatcherResponse{
				Watchers: []*api.Watcher{elm1},
			},
		},
		{
			name: "ok: len(watchers) >= page size + 1",
			req: &api.ListWatcherRequest{
				PageSize: 1,
			},
			mock: func(p *MockpersistentWatcher, _ *MocknameGenerator) {
				p.EXPECT().List(0, 2).Return([]*api.Watcher{elm1, elm2}, nil)
			},
			expected: &api.ListWatcherResponse{
				NextPageToken: "1",
				Watchers:      []*api.Watcher{elm1},
			},
		},
		{
			name: "invalid token",
			req: &api.ListWatcherRequest{
				PageToken: "abc",
			},
			code: codes.InvalidArgument,
		},
	}

	for _, x := range testcases {
		testWatcher(t, x.name, x.mock, x.code, x.expected, func(w *Watcher) (proto.Message, error) {
			return w.ListWatcher(context.TODO(), x.req)
		})
	}
}
