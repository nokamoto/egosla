package test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/egosla/api"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestState(t *testing.T) {
	s := make(State)
	key := "foo"

	expected := api.Watcher{
		Name: "bar",
	}

	s.Set(key, &expected)

	var actual api.Watcher
	if err := s.Get(key, &actual); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(&expected, &actual, protocmp.Transform()); len(diff) != 0 {
		t.Error(diff)
	}

	s.Delete(key)
	if err := s.Get(key, &actual); err == nil {
		t.Error("expected error but actual no error")
	}
}
