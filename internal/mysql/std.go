package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type std struct {
	db *gorm.DB
}

func (s *std) tx(f func(*gorm.DB) error) error {
	return s.txh("", f)
}

func (s *std) txh(hint string, f func(*gorm.DB) error) error {
	err := s.db.Transaction(f)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%w: %s", ErrNotFound, hint)
	}
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

func (s *std) get(name string, res interface{}) error {
	return s.txh(name, func(d *gorm.DB) error {
		return d.Where("name = ?", name).Take(res).Error
	})
}

func (s *std) list(offset, limit int, res interface{}) error {
	return s.tx(func(d *gorm.DB) error {
		return d.Offset(offset).Limit(limit).Find(res).Error
	})
}

func (s *std) update(name string, fields []string, updates interface{}, res interface{}) error {
	return s.txh(name, func(d *gorm.DB) error {
		if err := d.Model(updates).Where("name = ?", name).Select(fields).Updates(updates).Error; err != nil {
			return err
		}
		return d.Where("name = ?", name).Take(res).Error
	})
}
