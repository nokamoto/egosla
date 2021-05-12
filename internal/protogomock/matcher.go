package protogomock

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type matcher struct {
	x proto.Message
}

func (m matcher) Matches(x interface{}) bool {
	y, ok := x.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m.x, y)
}

func (m matcher) String() string {
	return fmt.Sprintf("equal(%v)", m.x)
}

// Equal is gomock.Matcher using proto.Equal.
func Equal(x proto.Message) matcher {
	return matcher{
		x: x,
	}
}
