package mysql

import (
	"errors"
	"fmt"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
)

type model interface {
	Typ() interface{}
	SliceTyp() interface{}
	Convert(proto.Message) (interface{}, error)
	Revert(interface{}) (proto.Message, error)
	RevertSlice(interface{}) ([]proto.Message, error)
	FieldMask(*field_mask.FieldMask) ([]string, error)
}

type WatcherModel struct{}

func (*WatcherModel) Typ() interface{} {
	return &watcher{}
}

func (*WatcherModel) SliceTyp() interface{} {
	return &[]watcher{}
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
	vs, ok := m.(*[]watcher)
	if !ok {
		return nil, fmt.Errorf("%s is not watcher slice", m)
	}
	var ms []proto.Message
	for _, v := range *vs {
		ms = append(ms, v.Value())
	}
	return ms, nil
}

func (*WatcherModel) FieldMask(updateMask *field_mask.FieldMask) ([]string, error) {
	updateMask.Normalize()
	var fields []string
	for _, path := range updateMask.GetPaths() {
		switch path {
		case "keywords":
			fields = append(fields, "keywords")
		default:
			return nil, fmt.Errorf("%s field not allowed", path)
		}
	}
	if len(fields) == 0 {
		return nil, errors.New("no update")
	}
	return fields, nil
}
