package mraccess_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mraccess"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	getter, err := mraccess.NewRolesGroupSet(testGroups(), testSource())
	require.NoError(t, err)

	id := [16]byte{1, 2, 3, 4}
	user := mraccess.NewUser(id, "administrators", "ru", getter)

	assert.Equal(t, id, user.ID())
	assert.Equal(t, "administrators", user.Group())
	assert.Equal(t, "ru", user.LangCode())
	assert.True(t, user.Has("users.manage"))
	assert.True(t, user.Has("public"))
	assert.False(t, user.Has("billing.manage"))
}

func TestNewUser_GroupRightsBinding(t *testing.T) {
	t.Parallel()

	getter, err := mraccess.NewRolesGroupSet(testGroups(), testSource())
	require.NoError(t, err)

	type testCase struct {
		name      string
		group     string
		right     string
		wantHas   bool
		wantGroup string
	}

	tests := []testCase{
		{"manager has orders", "managers", "orders.manage", true, "managers"},
		{"manager lacks users.manage", "managers", "users.manage", false, "managers"},
		{"guest has public", "guests", "public", true, "guests"},
		{"unknown group fail-closed", "ghost-group", "public", false, "ghost-group"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user := mraccess.NewUser([16]byte{}, tc.group, "en", getter)

			assert.Equal(t, tc.wantGroup, user.Group())
			assert.Equal(t, tc.wantHas, user.Has(tc.right))
		})
	}
}

func TestOneUserProvider(t *testing.T) {
	t.Parallel()

	getter, err := mraccess.NewRolesGroupSet(testGroups(), testSource())
	require.NoError(t, err)

	want := mraccess.NewUser([16]byte{9}, "guests", "en", getter)
	provider := mraccess.NewOneUserProvider(want)

	got, err := provider.UserByToken(t.Context(), "any-token")
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
