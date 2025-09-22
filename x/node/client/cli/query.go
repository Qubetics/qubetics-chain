package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	
	v3 "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"
	v1base "github.com/qubetics/qubetics-blockchain/v2/types/v1"
)

func queryNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node [node-addr]",
		Short: "Query a node by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := types.GetFromBech32(args[0], types.Bech32PrefixAccAddr)
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			res, err := qc.QueryNode(
				cmd.Context(),
				v3.NewQueryNodeRequest(addr),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Query all nodes with optional filters and pagination",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var id uint64
			if cmd.Flags().Changed("plan-id") {
				id, err = cmd.Flags().GetUint64("plan-id")
				if err != nil {
					return err
				}
			}

			var status v1base.Status
			if cmd.Flags().Changed("status") {
				statusStr, err := cmd.Flags().GetString("status")
				if err != nil {
					return err
				}
				status = v1base.StatusFromString(statusStr)
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			switch {
			case id != 0:
				res, err := qc.QueryNodesForPlan(
					cmd.Context(),
					v3.NewQueryNodesForPlanRequest(id, status, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			default:
				res, err := qc.QueryNodes(
					cmd.Context(),
					v3.NewQueryNodesRequest(status, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nodes")
	cmd.Flags().String("status", "", "filter the nodes by status (active|inactive)")
	cmd.Flags().Uint64("plan-id", 0, "filter the nodes by subscription plan ID")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-params",
		Short: "Query the node module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			res, err := qc.QueryParams(
				cmd.Context(),
				v3.NewQueryParamsRequest(),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
