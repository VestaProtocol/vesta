package cli

import (
	"strconv"

	"github.com/VestaProtocol/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [contract] [code]",
		Short: "Upgrade a smart contract to a new version",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContract := args[0]
			argCode := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpgrade(
				clientCtx.GetFromAddress().String(),
				argContract,
				argCode,
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
