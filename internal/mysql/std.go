package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

type std struct {
	db *gorm.DB
}

func (s *std) tx(f func(*gorm.DB) error) error {
	err := s.db.Transaction(f)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}
	return nil
}

func (s *std) create(created interface{}) error {
	return s.tx(func(d *gorm.DB) error {
		return d.Create(created).Error
	})
}
