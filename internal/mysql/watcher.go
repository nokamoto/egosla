package mysql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nokamoto/egosla/api"
	"gorm.io/gorm"
)

const (
	sep = ","
)

var (
	// ErrUnknown represents an error which is unneccessary to distinguish the cause or just unexpected.
	ErrUnknown = errors.New("unknown")
)

type watcher struct {
	Name string
	Keywords string
}

func newWatcher(v *api.Watcher) watcher {
	return watcher {
		Name: v.GetName(),
		Keywords: strings.Join(v.GetKeywords(), sep),
	}
}

func (w watcher) Value() *api.Watcher {
	return &api.Watcher{
		Name: w.Name,
		Keywords: strings.Split(w.Keywords, sep),
	}
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
func (p *PersistentWatcher)Create(v *api.Watcher) error {
	model := newWatcher(v)

	err := p.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(model).Error		
	})

	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	return nil
}


