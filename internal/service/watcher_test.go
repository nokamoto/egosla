package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/mysql"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func newTestWatcher(f func(*Watcher, proto.Message) (proto.Message, error)) stdCall {
	return func(l *zap.Logger, p persistent, ng nameGenerator, req proto.Message) (proto.Message, error) {
		return f(&Watcher{
			std: &std{
				logger:     l,
				persistent: p,
			},
			naming: ng,
		}, req)
	}
}

func TestWatcher_Create(t *testing.T) {
	req := &api.CreateWatcherRequest{
		Watcher: &api.Watcher{
			Keywords: []string{"bar"},
		},
	}

	res := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar"},
	}

	testcases := []stdTestCase{
		{
			name: "ok",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				mg.EXPECT().newName().Return(res.GetName())
				m.EXPECT().Create(prototest.Match(res)).Return(nil)
			},
			req: req,
			res: res,
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				mg.EXPECT().newName().Return(res.GetName())
				m.EXPECT().Create(gomock.Any()).Return(errors.New("unexpected"))
			},
			req:  &api.CreateWatcherRequest{},
			code: codes.Unavailable,
		},
	}

	testStd(
		t,
		newTestWatcher(func(w *Watcher, m proto.Message) (proto.Message, error) {
			return w.CreateWatcher(context.TODO(), m.(*api.CreateWatcherRequest))
		}),
		testcases,
	)
}

func TestWatcher_Get(t *testing.T) {
	req := &api.GetWatcherRequest{
		Name: "foo",
	}

	res := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar"},
	}

	testcases := []stdTestCase{
		{
			name: "ok",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Get(req.GetName()).Return(res, nil)
			},
			req: req,
			res: res,
		},
		{
			name: "not found",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Get(req.GetName()).Return(nil, mysql.ErrNotFound)
			},
			req:  req,
			code: codes.NotFound,
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Get(req.GetName()).Return(nil, errors.New("unexpected"))
			},
			req:  req,
			code: codes.Unavailable,
		},
	}

	testStd(
		t,
		newTestWatcher(func(w *Watcher, m proto.Message) (proto.Message, error) {
			return w.GetWatcher(context.TODO(), m.(*api.GetWatcherRequest))
		}),
		testcases,
	)
}

func TestWatcher_List(t *testing.T) {
	req := &api.ListWatcherRequest{
		PageSize: 1,
	}

	w1 := &api.Watcher{
		Name: "foo",
	}

	w2 := &api.Watcher{
		Name: "bar",
	}

	limit := int(req.GetPageSize() + 1)

	testcases := []stdTestCase{
		{
			name: "ok: empty",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().List(0, limit).Return(nil, nil)
			},
			req: req,
			res: &api.ListWatcherResponse{},
		},
		{
			name: "ok: <= size",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().List(0, limit).Return([]proto.Message{w1}, nil)
			},
			req: req,
			res: &api.ListWatcherResponse{
				Watchers: []*api.Watcher{w1},
			},
		},
		{
			name: "ok: > size",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().List(0, limit).Return([]proto.Message{w1, w2}, nil)
			},
			req: req,
			res: &api.ListWatcherResponse{
				NextPageToken: "1",
				Watchers:      []*api.Watcher{w1},
			},
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().List(0, limit).Return(nil, errors.New("unexpected"))
			},
			req:  req,
			code: codes.Unavailable,
		},
	}

	testStd(
		t,
		newTestWatcher(func(w *Watcher, m proto.Message) (proto.Message, error) {
			return w.ListWatcher(context.TODO(), m.(*api.ListWatcherRequest))
		}),
		testcases,
	)
}

func TestWatcher_Update(t *testing.T) {
	req := &api.UpdateWatcherRequest{
		Name: "foo",
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"keywords"},
		},
		Watcher: &api.Watcher{
			Keywords: []string{"bar"},
		},
	}

	res := &api.Watcher{
		Name:     "foo",
		Keywords: []string{"bar"},
	}

	testcases := []stdTestCase{
		{
			name: "ok",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().
					Update(req.GetName(), prototest.Match(req.GetUpdateMask()), prototest.Match(req.GetWatcher())).
					Return(res, nil)
			},
			req: req,
			res: res,
		},
		{
			name: "invalid argument",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, mysql.ErrInvalidArgument)
			},
			req:  req,
			code: codes.InvalidArgument,
		},
		{
			name: "not found",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, mysql.ErrNotFound)
			},
			req:  req,
			code: codes.NotFound,
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("unexpected"))
			},
			req:  req,
			code: codes.Unavailable,
		},
	}

	testStd(
		t,
		newTestWatcher(func(w *Watcher, m proto.Message) (proto.Message, error) {
			return w.UpdateWatcher(context.TODO(), m.(*api.UpdateWatcherRequest))
		}),
		testcases,
	)
}

func TestWatcherV2_Delete(t *testing.T) {
	req := &api.DeleteWatcherRequest{
		Name: "foo",
	}
	res := &empty.Empty{}
	testcases := []stdTestCase{
		{
			name: "ok",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Delete(req.GetName()).Return(nil)
			},
			req: req,
			res: res,
		},
		{
			name: "unexpected error",
			mock: func(m *Mockpersistent, mg *MocknameGenerator) {
				m.EXPECT().Delete(req.GetName()).Return(errors.New("unexpected"))
			},
			req:  req,
			code: codes.Unavailable,
		},
	}
	testStd(
		t,
		newTestWatcher(func(w *Watcher, m proto.Message) (proto.Message, error) {
			return w.DeleteWatcher(context.TODO(), m.(*api.DeleteWatcherRequest))
		}),
		testcases,
	)
}
