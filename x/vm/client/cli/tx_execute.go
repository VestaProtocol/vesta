package cli

import (
	"strconv"

	"vesta/x/vm/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdExecute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute [contract] [function] [args]",
		Short: "Execute a function from given contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContract := args[0]
			argFunction := args[1]
			argArgs := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExecute(
				clientCtx.GetFromAddress().String(),
				argContract,
				argFunction,
				argArgs,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
