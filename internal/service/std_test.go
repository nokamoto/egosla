package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
			ctrl := gomock.NewController(t)

			p := NewMockpersistent(ctrl)
			n := NewMocknameGenerator(ctrl)

			w := &WatcherV2{
				std: &std{
					logger:     zaptest.NewLogger(t),
					persistent: p,
				},
				naming: n,
			}

			if testcase.mock != nil {
				testcase.mock(p, n)
			}

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
