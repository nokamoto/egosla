//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"fmt"

	"github.com/google/uuid"
)

type nameGenerator interface{
	newName() string
}

type watcherNameGenerator struct{}

func (watcherNameGenerator) newName() string{
	return fmt.Sprintf("watchers/%s", uuid.NewString())
}
