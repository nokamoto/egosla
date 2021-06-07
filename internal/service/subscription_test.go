package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/prototest"
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
				p.EXPECT().Create(prototest.Match(expected)).Return(nil)
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
				p.EXPECT().Create(prototest.Match(expected)).Return(mysql.ErrUnknown)
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

func TestSubscription_List(t *testing.T) {
	elm1 := api.Subscription{
		Name: "foo",
	}

	elm2 := api.Subscription{
		Name: "bar",
	}

	testcases := []struct {
		name     string
		req      *api.ListSubscriptionRequest
		mock     func(p *MockpersistentSubscription, _ *MocknameGenerator)
		expected *api.ListSubscriptionResponse
		code     codes.Code
	}{
		{
			name: "ok",
			req: &api.ListSubscriptionRequest{
				PageSize: 1,
			},
			mock: func(p *MockpersistentSubscription, _ *MocknameGenerator) {
				p.EXPECT().List(0, 2).Return([]*api.Subscription{&elm1, &elm2}, nil)
			},
			expected: &api.ListSubscriptionResponse{
				NextPageToken: "1",
				Subscriptions: []*api.Subscription{&elm1},
			},
		},
		{
			name: "unexpected error",
			req:  &api.ListSubscriptionRequest{},
			mock: func(p *MockpersistentSubscription, _ *MocknameGenerator) {
				p.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, mysql.ErrUnknown)
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		testSubscription(t, x.name, x.mock, x.code, x.expected, func(s *Subscription) (proto.Message, error) {
			return s.ListSubscription(context.TODO(), x.req)
		})
	}
}

func TestSubscription_Delete(t *testing.T) {
	name := "foo"

	testcases := []struct {
		name     string
		mock     func(p *MockpersistentSubscription, n *MocknameGenerator)
		expected *empty.Empty
		code     codes.Code
	}{
		{
			name: "ok",
			mock: func(p *MockpersistentSubscription, n *MocknameGenerator) {
				p.EXPECT().Delete(name).Return(nil)
			},
			expected: &empty.Empty{},
		},
		{
			name: "unexpected error",
			mock: func(p *MockpersistentSubscription, n *MocknameGenerator) {
				p.EXPECT().Delete(name).Return(mysql.ErrUnknown)
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		testSubscription(t, x.name, x.mock, x.code, x.expected, func(s *Subscription) (proto.Message, error) {
			return s.DeleteSubscription(context.TODO(), &api.DeleteSubscriptionRequest{
				Name: name,
			})
		})
	}
}
