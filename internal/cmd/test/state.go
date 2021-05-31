package test

import (
	"fmt"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// State represents parameters passing to rest of scenarios.
type State map[string]string

// Set sets the protocol buffer message with the key.
func (s State) Set(key string, v proto.Message) {
	s[key] = prototext.Format(v)
}

// Get gets a protocol buffer message specified by the key.
func (s State) Get(key string, v proto.Message) error {
	str, ok := s[key]
	if !ok {
		return fmt.Errorf("state[%s] not found", key)
	}
	return prototext.Unmarshal([]byte(str), v)
}

// Delete deletes a protocol buffer message specified by the key.
func (s State) Delete(key string) {
	delete(s, key)
}
