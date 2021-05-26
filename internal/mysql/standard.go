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
