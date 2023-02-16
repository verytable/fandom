package titan

import "go.verytable.online/aot/aot/go/eren"

type AttackTitan struct {
}

func (t *AttackTitan) Abilities() []string {
	return eren.Abilities()
}
