package cli

import (
	"strconv"

	"github.com/VestaProtocol/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVmQuery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [name] [query] [args]",
		Short: "Query data from a smart contract.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqName := args[0]
			reqQuery := args[1]
			reqArgs := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQueryRequest{
				Name:  reqName,
				Query: reqQuery,
				Args:  reqArgs,
			}

			res, err := queryClient.Query(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
