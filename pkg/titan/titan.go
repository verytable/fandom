package titan

import (
	"go.verytable.online/aot/aot/go/eren"
	"go.verytable.online/aot/aot/go/mikasa"
)

type AttackTitan struct {
}

func (t *AttackTitan) Abilities() []string {
	return eren.Abilities()
}

func (t *AttackTitan) FriendAbilities() []string {
	return mikasa.Abilities()
}
