package mysql

import (
	"testing"

	"github.com/nokamoto/egosla/api"
	"github.com/nokamoto/egosla/internal/prototest"
)

func TestWatcherModel_Typ(t *testing.T) {
	m := &WatcherModel{}
	expected := &api.Watcher{}
	actual, err := m.Revert(m.Typ())
	if err != nil {
		t.Fatal(err)
	}
	if err := prototest.Equal(expected, actual); err != nil {
		t.Fatal(err)
	}
}

func TestWatcherModel_Convert(t *testing.T) {
	m := &WatcherModel{}

	expected := &api.Watcher{
		Name: "foo",
	}

	x, err := m.Convert(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := m.Revert(x)
	if err != nil {
		t.Fatal(err)
	}

	if err := prototest.Equal(expected, actual); err != nil {
		t.Fatal(err)
	}
}
