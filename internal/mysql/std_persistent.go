package mysql

import (
	"fmt"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type StdPersistent struct {
	std   *std
	model model
}

func NewStdPersistent(db *gorm.DB, model model) *StdPersistent {
	return &StdPersistent{
		std: &std{
			db: db,
		},
		model: model,
	}
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
	ms, err := s.model.RevertSlice(res)
	if err != nil {
		return nil, fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return ms, nil
}

func (s *StdPersistent) Update(name string, updateMask *field_mask.FieldMask, update proto.Message) (proto.Message, error) {
	v, err := s.model.Convert(update)
	if err != nil {
		return nil, fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	updateMask.Normalize()
	fields, err := s.model.FieldMask(updateMask)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidArgument, err)
	}
	res := s.model.Typ()
	if err := s.std.update(name, fields, v, res); err != nil {
		return nil, err
	}
	updated, err := s.model.Revert(res)
	if err != nil {
		return nil, fmt.Errorf("[bug] %w: %s", ErrUnknown, err)
	}
	return updated, nil
}

func (s *StdPersistent) Delete(name string) error {
	return s.std.delete(name, s.model.Typ())
}
