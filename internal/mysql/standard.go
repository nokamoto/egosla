package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

func createMethod(model interface{}, db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(model).Error
	})

	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	return nil
}

func listMethod(models interface{}, db *gorm.DB, offset, limit int) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		res := tx.Offset(offset).Limit(limit).Find(models)
		if res.Error != nil {
			return res.Error
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}
	return nil
}
