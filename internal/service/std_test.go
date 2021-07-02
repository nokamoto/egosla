package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestWatcherV2(t *testing.T, mock func(*Mockpersistent, *MocknameGenerator)) *WatcherV2 {
	ctrl := gomock.NewController(t)

	p := NewMockpersistent(ctrl)
	n := NewMocknameGenerator(ctrl)

	if mock != nil {
		mock(p, n)
	}

	return &WatcherV2{
		std: &std{
			logger:     zaptest.NewLogger(t),
			persistent: p,
		},
		naming: n,
	}
}

func TestWatcherV2_Create(t *testing.T) {
	requested := &api.CreateWatcherRequest{
		Watcher: &api.Watcher{
			Keywords: []string{"bar"},
		},
	}

	created := &api.Watcher{
		Name:     "foo",
		Keywords: requested.GetWatcher().GetKeywords(),
	}

	testcases := []struct {
		name string
		mock func(*Mockpersistent, *MocknameGenerator)
		req  *api.CreateWatcherRequest
		res  *api.Watcher
		code codes.Code
	}{
		{
			name: "ok",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				mg.EXPECT().newName().Return(created.GetName())
				m.EXPECT().Create(prototest.Match(created)).Return(nil)
			},
			req: requested,
			res: created,
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				mg.EXPECT().newName().Return(created.GetName())
				m.EXPECT().Create(prototest.Match(created)).Return(errors.New("unexpected"))
			},
			req:  requested,
			code: codes.Unavailable,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			w := newTestWatcherV2(t, testcase.mock)
			res, err := w.CreateWatcher(context.TODO(), testcase.req)
			if err := prototest.Equal(res, testcase.res); err != nil {
				t.Error(err)
			}
			if code := status.Code(err); code != testcase.code {
				t.Errorf("expected %v but actual %v", testcase.code, code)
			}
		})
	}
}

func TestWatcherV2_Get(t *testing.T) {
	requested := &api.GetWatcherRequest{
		Name: "foo",
	}

	got := &api.Watcher{
		Name:     requested.GetName(),
		Keywords: []string{"bar"},
	}

	testcases := []struct {
		name string
		mock func(*Mockpersistent, *MocknameGenerator)
		req  *api.GetWatcherRequest
		res  *api.Watcher
		code codes.Code
	}{
		{
			name: "ok",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Get(requested.GetName()).Return(got, nil)
			},
			req: requested,
			res: got,
		},
		{
			name: "not found",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Get(requested.GetName()).Return(nil, mysql.ErrNotFound)
			},
			req:  requested,
			code: codes.NotFound,
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Get(requested.GetName()).Return(nil, errors.New("unexpected"))
			},
			req:  requested,
			code: codes.Unavailable,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			w := newTestWatcherV2(t, testcase.mock)
			res, err := w.GetWatcher(context.TODO(), testcase.req)
			if err := prototest.Equal(res, testcase.res); err != nil {
				t.Error(err)
			}
			if code := status.Code(err); code != testcase.code {
				t.Errorf("expected %v but actual %v", testcase.code, code)
			}
		})
	}
}
