//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/moby/moby/pkg/namesgenerator"
)

type nameGenerator interface {
	newName() string
}

type watcherNameGenerator struct{}

func (watcherNameGenerator) newName() string {
	return fmt.Sprintf("watchers/%s", namesgenerator.GetRandomName(1))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
