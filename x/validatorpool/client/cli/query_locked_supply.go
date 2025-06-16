package cli

import (
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQueryLockedSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "locked-supply",
		Short: "Query the current locked supply in the validator pool",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLockedSupplyRequest{}

			res, err := queryClient.LockedSupply(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
