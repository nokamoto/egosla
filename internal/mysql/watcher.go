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

// Watcher represents a mysql record for api.Watcher.
type Watcher struct {
	Name string
	Keywords string
}

// NewWatcher creates a new Watcher from the api.Watcher.
func NewWatcher(v *api.Watcher) Watcher {
	return Watcher {
		Name: v.GetName(),
		Keywords: strings.Join(v.GetKeywords(), sep),
	}
}

// Value creates a new api.Watcher from the Watcher.
func (w Watcher) Value() *api.Watcher {
	return &api.Watcher{
		Name: w.Name,
		Keywords: strings.Split(w.Keywords, sep),
	}
}

// PersistentWatcher provides mysql operations for api.Watcher.
type PersistentWatcher struct {
	db *gorm.DB
}

// Create inserts the api.Watcher.
func (p *PersistentWatcher)Create(v *api.Watcher) error {
	model := NewWatcher(v)

	err := p.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(model).Error		
	})

	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	return nil
}


