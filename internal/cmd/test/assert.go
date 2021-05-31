package test

import (
	"errors"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

// Equal tests if x equals to y.
func Equal(x, y proto.Message) error {
	if diff := cmp.Diff(x, y, protocmp.Transform()); len(diff) != 0 {
		return errors.New(diff)
	}
	return nil
}
