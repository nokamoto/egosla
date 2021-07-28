package mysql

import (
	"errors"
	"fmt"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
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

type SubscriptionModel struct{}

func (*SubscriptionModel) Typ() interface{} {
	return &subscription{}
}

func (*SubscriptionModel) SliceTyp() interface{} {
	return &[]subscription{}
}

func (*SubscriptionModel) Convert(m proto.Message) (interface{}, error) {
	v, ok := m.(*api.Subscription)
	if !ok {
		return nil, fmt.Errorf("%s is not Watcher", m)
	}
	w := newSubscription(v)
	return &w, nil
}

func (*SubscriptionModel) Revert(m interface{}) (proto.Message, error) {
	v, ok := m.(*subscription)
	if !ok {
		return nil, fmt.Errorf("%s is not subscription", m)
	}
	return v.Value(), nil
}

func (*SubscriptionModel) RevertSlice(m interface{}) ([]proto.Message, error) {
	vs, ok := m.(*[]subscription)
	if !ok {
		return nil, fmt.Errorf("%s is not subscription slice", m)
	}
	var ms []proto.Message
	for _, v := range *vs {
		ms = append(ms, v.Value())
	}
	return ms, nil
}

func (*SubscriptionModel) FieldMask(updateMask *field_mask.FieldMask) ([]string, error) {
	var fields []string
	for _, path := range updateMask.GetPaths() {
		switch path {
		case "watcher":
			fields = append(fields, "watcher")
		default:
			return nil, fmt.Errorf("%s field not allowed", path)
		}
	}
	if len(fields) == 0 {
		return nil, errors.New("no update")
	}
	return fields, nil
}
