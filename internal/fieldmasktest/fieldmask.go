package fieldmasktest

import (
	"testing"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// NewValidFieldMask creates a valid FieldMask. It calls t.Fatal if paths is invalid.
func NewValidFieldMask(t *testing.T, v proto.Message, paths ...string) *field_mask.FieldMask {
	t.Helper()
	mask, err := fieldmaskpb.New(v, paths...)
	if err != nil {
		t.Fatal(err)
	}
	return mask
}
