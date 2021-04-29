package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestWatcher_Create(t *testing.T) {
	expected := &api.Watcher{
		Name: "watchers/foo",
		Keywords: []string{"bar", "baz"},
	}

	testcases := []struct{
		name string
		req *api.CreateWatcherRequest
		mock func(p *Mockpersistent, n *MocknameGenerator)
		expected *api.Watcher
		code codes.Code
	}{
		{
			name: "ok",
			req: &api.CreateWatcherRequest{
				Watcher: &api.Watcher{
					Keywords: expected.GetKeywords(),
				},
			},
			mock: func(p *Mockpersistent, n *MocknameGenerator) {
				n.EXPECT().newName().Return(expected.GetName())
				p.EXPECT().Create(expected).Return(nil)
			},
			expected: expected,
		},
		{
			name: "unexpected error",
			req: &api.CreateWatcherRequest{
				Watcher: &api.Watcher{
					Keywords: expected.GetKeywords(),
				},
			},
			mock: func(p *Mockpersistent, n *MocknameGenerator) {
				n.EXPECT().newName().Return(expected.GetName())
				p.EXPECT().Create(expected).Return(mysql.ErrUnknown)
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		t.Run(x.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			p := NewMockpersistent(ctrl)
			n := NewMocknameGenerator(ctrl)
			svc := &Watcher{
				p: p,
				n: n,
			}

			if x.mock != nil {
				x.mock(p, n)
			}

			actual, err := svc.CreateWatcher(context.TODO(), x.req)
			if code := status.Code(err); code != x.code {
				t.Errorf("expected %v but actual %v", x.code, code)
			}

			if diff := cmp.Diff(x.expected, actual, protocmp.Transform()); len(diff) != 0 {
				t.Error(diff)
			}
		})
	}
}
