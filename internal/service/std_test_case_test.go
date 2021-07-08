package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nokamoto/egosla/internal/prototest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type stdTestCase struct {
	name string
	mock func(*Mockpersistent, *MocknameGenerator)
	req  proto.Message
	res  proto.Message
	code codes.Code
}

type stdCall func(*zap.Logger, persistent, nameGenerator, proto.Message) (proto.Message, error)

func (s *stdTestCase) run(t *testing.T, f stdCall) {
	t.Run(s.name, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		p := NewMockpersistent(ctrl)
		n := NewMocknameGenerator(ctrl)
		if s.mock != nil {
			s.mock(p, n)
		}

		l := zaptest.NewLogger(t)
		res, err := f(l, p, n, s.req)
		if err := prototest.Equal(res, s.res); err != nil {
			t.Error(err)
		}
		if code := status.Code(err); code != s.code {
			t.Errorf("expected %v but actual %v", s.code, code)
		}
	})
}

func testStd(t *testing.T, f stdCall, xs []stdTestCase) {
	for _, x := range xs {
		x.run(t, f)
	}
}
