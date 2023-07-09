package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func makeNewAddress(address string) sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(address)
	return addr
}

func TestParseGuardianWhitelist_NonEmptyList(t *testing.T) {
	// Test case 1: Non-empty whitelist
	whitelist := "cosmos1yl6hdjhmkf37639730gffanpzndzdpmhwlkfhr,cosmos15phuqwakqlhpu7rjt8selspc2tmuzqf93uhp3g,cosmos1p3j465sr4svc5fylaz3aerlmk2fxw3umyd908k"
	expectedAddresses := []sdk.AccAddress{
		makeNewAddress("cosmos1yl6hdjhmkf37639730gffanpzndzdpmhwlkfhr"),
		makeNewAddress("cosmos15phuqwakqlhpu7rjt8selspc2tmuzqf93uhp3g"),
		makeNewAddress("cosmos1p3j465sr4svc5fylaz3aerlmk2fxw3umyd908k"),
	}

	res := ParseGuardianWhitelist(whitelist)
	// Cast result to []sdk.AccAddress
	castRes := make([]sdk.AccAddress, len(res))
	for i, addr := range res {
		castRes[i] = addr
	}

	require.Equal(t, expectedAddresses, castRes)
}

func TestParseGuardianWhitelist_EmptyList(t *testing.T) {
	// Test case 2: Empty whitelist
	emptyWhitelist := ""
	expectedEmptyAddresses := []sdk.AccAddress{}

	res := ParseGuardianWhitelist(emptyWhitelist)
	// Cast result to []sdk.AccAddress
	castRes := make([]sdk.AccAddress, len(res))
	for i, addr := range res {
		castRes[i] = addr
	}

	require.Equal(t, expectedEmptyAddresses, castRes)
}
