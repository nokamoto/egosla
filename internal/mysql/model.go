package mysql

import (
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
)

type model interface {
	Typ() interface{}
	SliceTyp() interface{}
	Convert(proto.Message) (interface{}, error)
	Revert(interface{}) (proto.Message, error)
	RevertSlice(interface{}) ([]proto.Message, error)
	FieldMask(*field_mask.FieldMask) ([]string, error)
}
