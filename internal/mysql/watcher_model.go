package mysql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
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
	k := strings.Split(w.Keywords, sep)
	if len(k) == 1 && k[0] == "" {
		k = nil
	}
	return &api.Watcher{
		Name:     w.Name,
		Keywords: k,
	}
}

func (watcher) TableName() string {
	return "watcher"
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
