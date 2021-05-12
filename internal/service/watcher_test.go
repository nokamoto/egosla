package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/fieldmasktest"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/protogomock"
	"go.uber.org/zap/zaptest"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

func mockWatcher(t *testing.T, f func(*Watcher, *Mockpersistent, *MocknameGenerator)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockpersistent(ctrl)
	n := NewMocknameGenerator(ctrl)
	svc := &Watcher{
		p:      p,
		n:      n,
		logger: zaptest.NewLogger(t),
	}

	f(svc, p, n)
}

func TestWatcher_Create(t *testing.T) {
	expected := &api.Watcher{
		Name:     "watchers/foo",
		Keywords: []string{"bar", "baz"},
	}

	testcases := []struct {
		name     string
		req      *api.CreateWatcherRequest
		mock     func(p *Mockpersistent, n *MocknameGenerator)
		expected *api.Watcher
		code     codes.Code
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
				p.EXPECT().Create(protogomock.Equal(expected)).Return(nil)
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
				p.EXPECT().Create(protogomock.Equal(expected)).Return(mysql.ErrUnknown)
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		t.Run(x.name, func(t *testing.T) {
			mockWatcher(t, func(svc *Watcher, p *Mockpersistent, n *MocknameGenerator) {
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
		})
	}
}

func TestWatcher_List(t *testing.T) {
	elm1 := &api.Watcher{
		Name: "foo",
	}

	elm2 := &api.Watcher{
		Name: "bar",
	}

	elm3 := &api.Watcher{
		Name: "baz",
	}

	testcases := []struct {
		name     string
		req      *api.ListWatcherRequest
		mock     func(p *Mockpersistent)
		expected *api.ListWatcherResponse
		code     codes.Code
	}{
		{
			name: "ok: empty token, empty watchers",
			req: &api.ListWatcherRequest{
				PageSize: 2,
			},
			mock: func(p *Mockpersistent) {
				p.EXPECT().List(0, 3).Return(nil, nil)
			},
			expected: &api.ListWatcherResponse{},
		},
		{
			name: "ok: len(watchers) < page size",
			req: &api.ListWatcherRequest{
				PageToken: "10",
				PageSize:  2,
			},
			mock: func(p *Mockpersistent) {
				p.EXPECT().List(10, 3).Return([]*api.Watcher{elm1}, nil)
			},
			expected: &api.ListWatcherResponse{
				Watchers: []*api.Watcher{elm1},
			},
		},
		{
			name: "ok: len(watchers) == page size",
			req: &api.ListWatcherRequest{
				PageToken: "10",
				PageSize:  2,
			},
			mock: func(p *Mockpersistent) {
				p.EXPECT().List(10, 3).Return([]*api.Watcher{elm1, elm2}, nil)
			},
			expected: &api.ListWatcherResponse{
				Watchers: []*api.Watcher{elm1, elm2},
			},
		},
		{
			name: "ok: len(watchers) >= page size + 1",
			req: &api.ListWatcherRequest{
				PageToken: "10",
				PageSize:  2,
			},
			mock: func(p *Mockpersistent) {
				p.EXPECT().List(10, 3).Return([]*api.Watcher{elm1, elm2, elm3}, nil)
			},
			expected: &api.ListWatcherResponse{
				NextPageToken: "12",
				Watchers:      []*api.Watcher{elm1, elm2},
			},
		},
		{
			name: "unexpected error",
			req: &api.ListWatcherRequest{
				PageSize: 2,
			},
			mock: func(p *Mockpersistent) {
				p.EXPECT().List(0, 3).Return(nil, mysql.ErrUnknown)
			},
			code: codes.Unavailable,
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
		t.Run(x.name, func(t *testing.T) {
			mockWatcher(t, func(svc *Watcher, p *Mockpersistent, n *MocknameGenerator) {
				if x.mock != nil {
					x.mock(p)
				}

				actual, err := svc.ListWatcher(context.TODO(), x.req)
				if code := status.Code(err); code != x.code {
					t.Errorf("expected %v but actual %v", x.code, code)
				}

				if diff := cmp.Diff(x.expected, actual, protocmp.Transform()); len(diff) != 0 {
					t.Error(diff)
				}
			})
		})
	}
}

func TestWatcher_Delete(t *testing.T) {
	name := "foo"

	testcases := []struct {
		name string
		mock func(p *Mockpersistent, n *MocknameGenerator)
		code codes.Code
	}{
		{
			name: "ok",
			mock: func(p *Mockpersistent, n *MocknameGenerator) {
				p.EXPECT().Delete(name).Return(nil)
			},
		},
		{
			name: "unexpected error",
			mock: func(p *Mockpersistent, n *MocknameGenerator) {
				p.EXPECT().Delete(name).Return(mysql.ErrUnknown)
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		t.Run(x.name, func(t *testing.T) {
			mockWatcher(t, func(svc *Watcher, p *Mockpersistent, n *MocknameGenerator) {
				if x.mock != nil {
					x.mock(p, n)
				}

				_, err := svc.DeleteWatcher(context.TODO(), &api.DeleteWatcherRequest{
					Name: name,
				})
				if code := status.Code(err); code != x.code {
					t.Errorf("expected %v but actual %v", x.code, code)
				}
			})
		})
	}
}

func TestWatcher_Update(t *testing.T) {
	watcher := &api.Watcher{
		Keywords: []string{"bar"},
	}

	result := &api.Watcher{
		Name:     "foo",
		Keywords: watcher.GetKeywords(),
	}

	updateMask := fieldmasktest.NewValidFieldMask(t, watcher, "keywords")

	testcases := []struct {
		name     string
		req      *api.UpdateWatcherRequest
		mock     func(p *Mockpersistent, n *MocknameGenerator)
		expected *api.Watcher
		code     codes.Code
	}{
		{
			name: "ok",
			req: &api.UpdateWatcherRequest{
				Name:       result.GetName(),
				Watcher:    watcher,
				UpdateMask: updateMask,
			},
			mock: func(p *Mockpersistent, n *MocknameGenerator) {
				p.EXPECT().Update(protogomock.Equal(result), protogomock.Equal(updateMask)).Return(result, nil)
			},
			expected: result,
		},
		{
			name: "err: name contains",
			req: &api.UpdateWatcherRequest{
				Name:       result.GetName(),
				Watcher:    watcher,
				UpdateMask: fieldmasktest.NewValidFieldMask(t, watcher, "name", "keywords"),
			},
			code: codes.InvalidArgument,
		},
		{
			name: "err: unknown fields",
			req: &api.UpdateWatcherRequest{
				Name:    result.GetName(),
				Watcher: watcher,
				UpdateMask: &field_mask.FieldMask{
					Paths: []string{"foo"},
				},
			},
			code: codes.InvalidArgument,
		},
		{
			name: "unexpected error",
			req: &api.UpdateWatcherRequest{
				Name:       result.GetName(),
				Watcher:    watcher,
				UpdateMask: updateMask,
			},
			mock: func(p *Mockpersistent, n *MocknameGenerator) {
				p.EXPECT().Update(protogomock.Equal(result), protogomock.Equal(updateMask)).Return(nil, errors.New("unexpected"))
			},
			code: codes.Unavailable,
		},
	}

	for _, x := range testcases {
		t.Run(x.name, func(t *testing.T) {
			mockWatcher(t, func(svc *Watcher, p *Mockpersistent, n *MocknameGenerator) {
				if x.mock != nil {
					x.mock(p, n)
				}

				actual, err := svc.UpdateWatcher(context.TODO(), x.req)
				if code := status.Code(err); code != x.code {
					t.Errorf("expected %v but actual %v", x.code, code)
				}

				if diff := cmp.Diff(x.expected, actual, protocmp.Transform()); len(diff) != 0 {
					t.Error(diff)
				}
			})
		})
	}
}
