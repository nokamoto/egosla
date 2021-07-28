package mysql

import (
	"testing"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/proto"
)

func TestSubscriptionModel(t *testing.T) {
	testModel(t, modelTestSet{
		m:        &SubscriptionModel{},
		typValue: &api.Subscription{},
		nonZeroValue: &api.Subscription{
			Name: "foo",
		},
		mysqlValues: &[]subscription{
			{
				Watcher: "foo",
			},
			{
				Watcher: "bar",
			},
		},
		protoValues: []proto.Message{
			&api.Subscription{
				Watcher: "foo",
			},
			&api.Subscription{
				Watcher: "bar",
			},
		},
		mask: map[string]string{
			"watcher": "watcher",
		},
	})
}
