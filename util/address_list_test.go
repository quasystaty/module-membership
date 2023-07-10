package util

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestSplitStringIntoAddresses(t *testing.T) {
	valid_1 := "cosmos1l0znsvddllw9knha3yx2svnlxny676d8ns7uys"
	valid_2 := "cosmos1j8pp7zvcu9z8vd882m284j29fn2dszh05cqvf9"
	invalid := "invalidaddress"

	t.Run("Valid addresses", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses(fmt.Sprintf("%s,%s", valid_1, valid_2))
		assert.NoError(t, err)
		assert.Lenf(t, addresses, 2, "Expected 2 addresses, got: %d", len(addresses))

		for _, addr := range addresses {
			if _, err := types.AccAddressFromBech32(addr.String()); err != nil {
				assert.Error(t, err, "Expected valid AccAddress")
			}
		}
	})

	t.Run("Invalid addresses", func(t *testing.T) {
		_, err := SplitStringIntoAddresses(fmt.Sprintf("%s_1,%s_2", invalid, invalid))
		assert.Error(t, err)
		assert.ErrorContains(t, err, "decoding bech32 failed")
	})

	t.Run("Empty string", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses("")
		assert.ErrorContains(t, err, "empty address")
		assert.Nil(t, addresses)
	})

	t.Run("Single address", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses(valid_1)
		assert.NoError(t, err)
		assert.Len(t, addresses, 1, "Expected 1 address")
		if _, err := types.AccAddressFromBech32(addresses[0].String()); err != nil {
			assert.Errorf(t, err, "Expected valid AccAddress")
		}
	})

	t.Run("Address with leading/trailing whitespace", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses(fmt.Sprintf(" %s , %s ", valid_1, valid_2))
		assert.NoError(t, err)
		assert.Len(t, addresses, 2, "Expected 2 addresses")
		for _, addr := range addresses {
			if _, err := types.AccAddressFromBech32(addr.String()); err != nil {
				assert.NoErrorf(t, err, "Expected valid AccAddress")
			}
		}
	})

	t.Run("Address with extra commas", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses(fmt.Sprintf(",%s,,%s,", valid_1, valid_2))
		assert.ErrorContains(t, err, "empty address")
		assert.Nil(t, addresses)
	})
}
