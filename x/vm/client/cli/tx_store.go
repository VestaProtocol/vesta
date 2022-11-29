package cli

import (
	"os"
	"strconv"

	"github.com/VestaProtocol/vesta/x/vm/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [source]",
		Short: "Storing source code to the chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSource := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			b, err := os.ReadFile(argSource) // just pass the file name
			if err != nil {
				return err
			}

			codeStr := string(b) // convert content to a 'string'

			msg := types.NewMsgStore(
				clientCtx.GetFromAddress().String(),
				codeStr,
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
