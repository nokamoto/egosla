//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"github.com/nokamoto/egosla/api"
)

type persistent interface {
	Create(*api.Watcher) error
	List(offset, limit int) ([]*api.Watcher, error)
}
