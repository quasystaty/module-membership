package util

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
)

func TestSplitStringIntoAddresses(t *testing.T) {
	validAddresses := "cosmos1abcdefgh1234567890,cosmos1ijklmnopqrstuvwxyz0987654321"
	invalidAddresses := "cosmos1abcdefgh1234567890,invalidaddress"

	t.Run("Valid addresses", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses(validAddresses)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if len(addresses) != 2 {
			t.Errorf("Expected 2 addresses, got: %d", len(addresses))
		}

		for _, addr := range addresses {
			if _, err := types.AccAddressFromBech32(addr.String()); err != nil {
				t.Errorf("Expected valid AccAddress, got: %v", err)
			}
		}
	})

	t.Run("Invalid addresses", func(t *testing.T) {
		_, err := SplitStringIntoAddresses(invalidAddresses)
		if err == nil {
			t.Error("Expected an error, got nil")
		}

		expectedErrMsg := "invalid addresses"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', got: '%s'", expectedErrMsg, err.Error())
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses("")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if len(addresses) != 0 {
			t.Errorf("Expected 0 addresses, got: %d", len(addresses))
		}
	})

	t.Run("Single address", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses("cosmos1abcdefgh1234567890")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if len(addresses) != 1 {
			t.Errorf("Expected 1 address, got: %d", len(addresses))
		}

		if _, err := types.AccAddressFromBech32(addresses[0].String()); err != nil {
			t.Errorf("Expected valid AccAddress, got: %v", err)
		}
	})

	t.Run("Address with leading/trailing whitespace", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses(" cosmos1abcdefgh1234567890 , cosmos1ijklmnopqrstuvwxyz0987654321 ")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if len(addresses) != 2 {
			t.Errorf("Expected 2 addresses, got: %d", len(addresses))
		}

		for _, addr := range addresses {
			if _, err := types.AccAddressFromBech32(addr.String()); err != nil {
				t.Errorf("Expected valid AccAddress, got: %v", err)
			}
		}
	})

	t.Run("Address with extra commas", func(t *testing.T) {
		addresses, err := SplitStringIntoAddresses("cosmos1abcdefgh1234567890,,,cosmos1ijklmnopqrstuvwxyz0987654321")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if len(addresses) != 2 {
			t.Errorf("Expected 2 addresses, got: %d", len(addresses))
		}

		for _, addr := range addresses {
			if _, err := types.AccAddressFromBech32(addr.String()); err != nil {
				t.Errorf("Expected valid AccAddress, got: %v", err)
			}
		}
	})
}
