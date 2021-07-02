package mysql

import (
	"fmt"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/proto"
)

type model interface {
	Typ() interface{}
	SliceTyp() interface{}
	Convert(proto.Message) (interface{}, error)
	Revert(interface{}) (proto.Message, error)
	RevertSlice(interface{}) ([]proto.Message, error)
}

type WatcherModel struct{}

func (*WatcherModel) Typ() interface{} {
	return &watcher{}
}

func (*WatcherModel) SliceTyp() interface{} {
	return []watcher{}
}

func (*WatcherModel) Convert(m proto.Message) (interface{}, error) {
	v, ok := m.(*api.Watcher)
	if !ok {
		return nil, fmt.Errorf("%s is not Watcher", m)
	}
	w := newWatcher(v)
	return &w, nil
}

func (*WatcherModel) Revert(m interface{}) (proto.Message, error) {
	v, ok := m.(*watcher)
	if !ok {
		return nil, fmt.Errorf("%s is not watcher", m)
	}
	return v.Value(), nil
}

func (*WatcherModel) RevertSlice(m interface{}) ([]proto.Message, error) {
	vs, ok := m.([]watcher)
	if !ok {
		return nil, fmt.Errorf("%s is not watcher slice", m)
	}
	var ms []proto.Message
	for _, v := range vs {
		ms = append(ms, v.Value())
	}
	return ms, nil
}
