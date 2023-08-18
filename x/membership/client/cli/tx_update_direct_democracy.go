package cli

import (
	"fmt"
	"strconv"

	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const (
	FlagAddGuardians      = "add-guardians"
	FlagRemoveGuardians   = "remove-guardians"
	FlagTotalVotingWeight = "total-voting-weight"
)

func NewSubmitUpdateDirectDemocracyProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-direct-democracy",
		Short: "Submit a proposal to update the Direct Democracy settings",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal to add or remove guardians from the electorate, 
or to update the total voting weight of the electorate. The proposal will be put to a governance vote
and only applied if it succeeds.

NOTE: At least one change must be submitted using the flags below. Updates may be combined by using multiple flags.

Example - Adding a guardian:
$ %s tx membership update-direct-democracy --add-guardians=<new_guardian> --from=<key_or_address>

Example - Removing a guardian:
$ %s tx membership update-direct-democracy --remove-guardians=<guardian_to_remove> --from=<key_or_address>

Example - Updating the total voting weight:
$ %s tx membership update-direct-democracy --total-voting-weight=1000 --from=<key_or_address>

`,
				version.AppName,
				version.AppName,
				version.AppName,
			)),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			addGuardians, err := cmd.Flags().GetStringArray(FlagAddGuardians)
			if err != nil {
				return err
			}

			removeGuardians, err := cmd.Flags().GetStringArray(FlagRemoveGuardians)
			if err != nil {
				return err
			}

			totalVotingWeight, err := cmd.Flags().GetString(FlagTotalVotingWeight)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDirectDemocracy(
				clientCtx.GetFromAddress().String(),
				addGuardians,
				removeGuardians,
				totalVotingWeight,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Add flags
	cmd.Flags().String(FlagTotalVotingWeight, "", "total voting weight of the electorate")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
