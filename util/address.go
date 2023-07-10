package util

import (
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SplitStringIntoAddresses(str string) ([]sdk.AccAddress, error) {
	var addresses = []sdk.AccAddress{}

	for _, addr := range strings.Split(str, ",") {
		trimmedAddr := strings.TrimSpace(addr)
		if len(trimmedAddr) == 0 {
			return nil, errors.New("empty address")
		}
		validAddr, err := sdk.AccAddressFromBech32(trimmedAddr)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, validAddr)
	}
	return addresses, nil
}
