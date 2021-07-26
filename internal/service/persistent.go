//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
)

type persistent interface {
	Create(proto.Message) error
	Get(string) (proto.Message, error)
	List(int, int) ([]proto.Message, error)
	Update(string, *field_mask.FieldMask, proto.Message) (proto.Message, error)
	Delete(string) error
}
