package mysql

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type StdPersistent struct {
	std     *std
	convert func(proto.Message) (interface{}, error)
}

func (s *StdPersistent) Create(created proto.Message) error {
	v, err := s.convert(created)
	if err != nil {
		return fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return s.std.create(v)
}
