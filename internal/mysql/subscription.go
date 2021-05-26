package mysql

import (
	"github.com/nokamoto/egosla/api"
	"gorm.io/gorm"
)

type subscription struct {
	Name    string
	Watcher string
}

func newSubscription(v *api.Subscription) subscription {
	return subscription{
		Name:    v.GetName(),
		Watcher: v.GetWatcher(),
	}
}

func (s subscription) Value() *api.Subscription {
	return &api.Subscription{
		Name:    s.Name,
		Watcher: s.Watcher,
	}
}

func (subscription) TableName() string {
	return "subscription"
}

// PersistentSubscription provides mysql operations for api.Subscription.
type PersistentSubscription struct {
	db *gorm.DB
}

// NewPersistentSubscription creates a new PersistentSubscription.
func NewPersistentSubscriptionn(db *gorm.DB) *PersistentSubscription {
	return &PersistentSubscription{
		db: db,
	}
}

// Create inserts the api.Watcher.
func (p *PersistentSubscription) Create(v *api.Subscription) error {
	return createMethod(newSubscription(v), p.db)
}
