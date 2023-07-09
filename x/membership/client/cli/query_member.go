package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/spf13/cobra"
)

const (
	FlagAddress = "address"
)

var _ = strconv.Itoa(0)

func CmdMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member [address]",
		Short: "Query a member's account using their wallet address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			address := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMemberRequest{
				Address: address,
			}

			res, err := queryClient.Member(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
