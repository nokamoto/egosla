package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/protogomock"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func testSubscription(
	t *testing.T,
	name string,
	mock func(p *MockpersistentSubscription, n *MocknameGenerator),
	expectedCode codes.Code,
	expectedResponse proto.Message,
	call func(*Subscription) (proto.Message, error),
) {
	t.Run(name, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		p := NewMockpersistentSubscription(ctrl)
		n := NewMocknameGenerator(ctrl)
		svc := &Subscription{
			p:      p,
			n:      n,
			logger: zaptest.NewLogger(t),
		}

		if mock != nil {
			mock(p, n)
		}

		res, err := call(svc)
		if code := status.Code(err); code != expectedCode {
			t.Errorf("expected %v but actual %v", expectedCode, code)
		}

		if diff := cmp.Diff(expectedResponse, res, protocmp.Transform()); len(diff) != 0 {
			t.Error(diff)
		}
	})
}

func TestSubscription_Create(t *testing.T) {
	expected := &api.Subscription{
		Name:    "subscriptions/foo",
		Watcher: "watchers/bar",
	}

	testcases := []struct {
		name     string
		req      *api.CreateSubscriptionRequest
		mock     func(p *MockpersistentSubscription, n *MocknameGenerator)
		expected *api.Subscription
		code     codes.Code
	}{
		{
			name: "ok",
			req: &api.CreateSubscriptionRequest{
				Subscription: &api.Subscription{
					Watcher: expected.GetWatcher(),
				},
			},
			mock: func(p *MockpersistentSubscription, n *MocknameGenerator) {
				n.EXPECT().newName().Return(expected.GetName())
				p.EXPECT().Create(protogomock.Equal(expected)).Return(nil)
			},
			expected: expected,
		},
		{
			name: "unexpected error",
			req: &api.CreateSubscriptionRequest{
				Subscription: &api.Subscription{
					Watcher: expected.GetWatcher(),
				},
			},
			mock: func(p *MockpersistentSubscription, n *MocknameGenerator) {
				n.EXPECT().newName().Return(expected.GetName())
				p.EXPECT().Create(protogomock.Equal(expected)).Return(mysql.ErrUnknown)
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		testSubscription(t, x.name, x.mock, x.code, x.expected, func(s *Subscription) (proto.Message, error) {
			return s.CreateSubscription(context.TODO(), x.req)
		})
	}
}