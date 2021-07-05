package prototest

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

// Equal tests if x equals to y.
func Equal(x, y proto.Message) error {
	if (x == nil || reflect.ValueOf(x).IsNil()) && (y == nil || reflect.ValueOf(y).IsNil()) {
		return nil
	}
	if diff := cmp.Diff(x, y, protocmp.Transform()); len(diff) != 0 {
		return errors.New(diff)
	}
	return nil
}

func EqualSlice(x, y []proto.Message) error {
	if len(x) != len(y) {
		return fmt.Errorf("len(%v) != len(%v)", x, y)
	}
	for i, z := range x {
		if err := Equal(z, y[i]); err != nil {
			return err
		}
	}
	return nil
}
