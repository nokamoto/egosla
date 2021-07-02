package mysql

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

type StdPersistent struct {
	std   *std
	model model
}

func (s *StdPersistent) Create(created proto.Message) error {
	v, err := s.model.Convert(created)
	if err != nil {
		return fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return s.std.create(v)
}

func (s *StdPersistent) Get(name string) (proto.Message, error) {
	res := s.model.Typ()
	if err := s.std.get(name, res); err != nil {
		return nil, err
	}
	v, err := s.model.Revert(res)
	if err != nil {
		return nil, fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return v, nil
}

func (s *StdPersistent) List(offset, limit int) ([]proto.Message, error) {
	res := s.model.SliceTyp()
	if err := s.std.list(offset, limit, res); err != nil {
		return nil, err
	}
	log.Printf("res=%v", res)
	ms, err := s.model.RevertSlice(res)
	if err != nil {
		return nil, fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return ms, nil
}
