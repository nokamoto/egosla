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

type prefixNameGenerator struct {
	prefix string
}

func (p prefixNameGenerator) newName() string {
	return fmt.Sprintf("%s/%s", p.prefix, namesgenerator.GetRandomName(1))
}

func newWatcherNameGenerator() prefixNameGenerator {
	return prefixNameGenerator{
		prefix: "watchers",
	}
}

func newSubscriptionNameGenerator() prefixNameGenerator {
	return prefixNameGenerator{
		prefix: "subscriptions",
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
