package mysql

import (
	"testing"

	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/proto"
)

func TestWatcherModel(t *testing.T) {
	testModel(t, modelTestSet{
		m:            &WatcherModel{},
		typValue:     &api.Watcher{},
		nonZeroValue: &api.Watcher{Name: "foo"},
		mysqlValues: &[]watcher{
			{
				Name: "foo",
			},
			{
				Name: "bar",
			},
		},
		protoValues: []proto.Message{
			&api.Watcher{
				Name: "foo",
			},
			&api.Watcher{
				Name: "bar",
			},
		},
		mask: map[string]string{
			"keywords": "keywords",
		},
	})
}
