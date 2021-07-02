package mysql

import (
	"testing"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/prototest"
	"google.golang.org/protobuf/proto"
)

type modelTestSet struct {
	m            model
	typValue     proto.Message
	nonZeroValue proto.Message
	mysqlValues  interface{}
	protoValues  []proto.Message
}

func testModel(t *testing.T, set modelTestSet) {
	t.Helper()

	t.Run("Revert(Typ) should be typValue", func(t *testing.T) {
		expected := set.typValue
		actual, err := set.m.Revert(set.m.Typ())
		if err != nil {
			t.Fatal(err)
		}
		if err := prototest.Equal(expected, actual); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Revert(Convert) should preserve nonZeroValue", func(t *testing.T) {
		expected := set.nonZeroValue
		converted, err := set.m.Convert(expected)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := set.m.Revert(converted)
		if err != nil {
			t.Fatal(err)
		}
		if err := prototest.Equal(expected, actual); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("RevertSlice(SliceTyp) should be nil", func(t *testing.T) {
		actual, err := set.m.RevertSlice(set.m.SliceTyp())
		if err != nil {
			t.Fatal(err)
		}
		if err := prototest.EqualSlice(nil, actual); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("RevertSlice(SliceTyp) should be nil", func(t *testing.T) {
		actual, err := set.m.RevertSlice(set.m.SliceTyp())
		if err != nil {
			t.Fatal(err)
		}
		if err := prototest.EqualSlice(nil, actual); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("RevertSlice(mysqlValues) should be protoValues", func(t *testing.T) {
		expected := set.protoValues
		actual, err := set.m.RevertSlice(set.mysqlValues)
		if err != nil {
			t.Fatal(err)
		}
		if err := prototest.EqualSlice(expected, actual); err != nil {
			t.Fatal(err)
		}
	})
}

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
	})
}
