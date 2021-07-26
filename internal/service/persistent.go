//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"github.com/nokamoto/egosla/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
)

type persistentWatcher interface {
	Create(*api.Watcher) error
	List(offset, limit int) ([]*api.Watcher, error)
	Delete(name string) error
	Update(*api.Watcher, *field_mask.FieldMask) (*api.Watcher, error)
	Get(name string) (*api.Watcher, error)
}

type persistentSubscription interface {
	Create(*api.Subscription) error
	List(offset, limit int) ([]*api.Subscription, error)
	Delete(name string) error
	Get(name string) (*api.Subscription, error)
}

type persistent interface {
	Create(proto.Message) error
	Get(string) (proto.Message, error)
	List(int, int) ([]proto.Message, error)
	Update(string, *field_mask.FieldMask, proto.Message) (proto.Message, error)
	Delete(string) error
}
