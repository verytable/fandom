package titan_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.verytable.online/fandom/pkg/titan"
)

func TestAttackTitan_Abilities(t *testing.T) {
	titan := &titan.AttackTitan{}
	require.Contains(t, titan.Abilities(), "Attack Titan")
}

func TestAttackTitan_FriendAbilities(t *testing.T) {
	titan := &titan.AttackTitan{}
	require.Contains(t, titan.FriendAbilities(), "Strength")
}
