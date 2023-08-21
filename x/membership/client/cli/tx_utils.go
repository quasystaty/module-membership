package cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

const (
	flagDeposit = "deposit"
)

func parseInitialDeposit(cmd *cobra.Command) (sdk.Coins, error) {
	// parse initial deposit
	depositStr, err := cmd.Flags().GetString(flagDeposit)
	if err != nil {
		return nil, fmt.Errorf("no initial deposit found. did you set --deposit? %s", err)
	}
	deposit, err := sdk.ParseCoinsNormalized(depositStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse deposit: %s", err)
	}
	if !deposit.IsValid() || deposit.IsZero() {
		return nil, fmt.Errorf("no initial deposit set, use --deposit flag")
	}
	return deposit, nil
}
