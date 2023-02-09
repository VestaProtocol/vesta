package keeper

import (
	"fmt"

	sdkErrors "cosmossdk.io/errors"
	"github.com/VestaProtocol/vesta/x/funding/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmosErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) HandleTokens(ctx sdk.Context) {
	ctx.Logger().Error("Handling tokens!")
	senderAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if senderAddr == nil {
		panic(sdkErrors.Wrapf(cosmosErrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	balances := k.bankKeeper.GetAllBalances(ctx, senderAddr)

	modifier := sdk.NewInt(100) // we take whatever the balance divided by this number is each block.

	all := []sdk.Coin{}
	for _, coin := range balances {
		ctx.Logger().Error(fmt.Sprintf("Adding %d %s", coin.Amount.Int64(), coin.Denom))

		a := coin.Amount.Quo(modifier)

		c := sdk.NewCoin(coin.Denom, a)
		ctx.Logger().Error(fmt.Sprintf("Finishing with %d %s", c.Amount.Int64(), c.Denom))

		all = append(all, c)
	}
	fees := sdk.NewCoins(all...)
	ctx.Logger().Error(fmt.Sprintf("Total fees sent: %s", fees.String()))

	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
	if err != nil {
		panic(err)
	}
}
