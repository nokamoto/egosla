package mysql

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type StdPersistent struct {
	std     *std
	convert func(proto.Message) (interface{}, error)
	typ     func() interface{}
	revert  func(interface{}) (proto.Message, error)
}

func (s *StdPersistent) Create(created proto.Message) error {
	v, err := s.convert(created)
	if err != nil {
		return fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return s.std.create(v)
}

func (s *StdPersistent) Get(name string) (proto.Message, error) {
	res := s.typ()
	if err := s.std.get(name, res); err != nil {
		return nil, err
	}
	v, err := s.revert(res)
	if err != nil {
		return nil, fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return v, nil
}
