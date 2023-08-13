package types_test

import (
	"testing"

	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/require"
)

func TestAllKeys_NoDuplicates(t *testing.T) {
	// loop through AllKeys and ensure there are no duplicate keys
	seen := make(map[string]bool)
	for _, key := range types.AllKeys {
		// Convert key to a string
		s := string(key)
		require.False(t, seen[s], "duplicate key: %s", key)
		seen[s] = true
	}
}
