package mysql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
)

const (
	sep = ","
)

type watcher struct {
	Name     string
	Keywords string
}

func newWatcher(v *api.Watcher) watcher {
	return watcher{
		Name:     v.GetName(),
		Keywords: strings.Join(v.GetKeywords(), sep),
	}
}

func (w watcher) Value() *api.Watcher {
	return &api.Watcher{
		Name:     w.Name,
		Keywords: strings.Split(w.Keywords, sep),
	}
}

func (watcher) TableName() string {
	return "watcher"
}

// PersistentWatcher provides mysql operations for api.Watcher.
type PersistentWatcher struct {
	db *gorm.DB
}

// NewPersistentWatcher creates a new PersistentWatcher.
func NewPersistentWatcher(db *gorm.DB) *PersistentWatcher {
	return &PersistentWatcher{
		db: db,
	}
}

// Create inserts the api.Watcher.
func (p *PersistentWatcher) Create(v *api.Watcher) error {
	return createMethod(newWatcher(v), p.db)
}

// List selects a list of watchers from offset to limit.
func (p *PersistentWatcher) List(offset, limit int) ([]*api.Watcher, error) {
	var watchers []*api.Watcher
	err := p.db.Transaction(func(tx *gorm.DB) error {
		var ms []watcher
		res := tx.Offset(offset).Limit(limit).Find(&ms)
		if res.Error != nil {
			return res.Error
		}

		for _, m := range ms {
			watchers = append(watchers, m.Value())
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	return watchers, nil
}

// Delete deletes a watcher by the name.
func (p *PersistentWatcher) Delete(name string) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("name = ?", name).Delete(&watcher{})
		return res.Error
	})

	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	return nil
}

// Update updates the watcher by the name with the update mask.
func (p *PersistentWatcher) Update(v *api.Watcher, updateMask *field_mask.FieldMask) (*api.Watcher, error) {
	var updated watcher

	updateMask.Normalize()
	var fields []string
	for _, path := range updateMask.GetPaths() {
		switch path {
		case "name":
			return nil, fmt.Errorf("%w: name field not allowed", ErrInvalidArgument)
		case "keywords":
			fields = append(fields, "keywords")
		}
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("%w: no update", ErrInvalidArgument)
	}

	err := p.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&watcher{}).Where("name = ?", v.GetName()).Select(fields).Updates(newWatcher(v))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return fmt.Errorf("expected 1 but actual %v", res.RowsAffected)
		}

		res = tx.First(&updated, "name = ?", v.GetName())
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return fmt.Errorf("expected 1 but actual %v", res.RowsAffected)
		}
		return nil
	})

	if errors.Is(err, ErrInvalidArgument) {
		return nil, err
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	return updated.Value(), nil
}
