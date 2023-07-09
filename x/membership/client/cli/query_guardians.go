package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGuardians() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "guardians",
		Short: "Query the guardians of the electorate",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGuardiansRequest{}

			res, err := queryClient.Guardians(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
