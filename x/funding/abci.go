package funding

import (
	"github.com/VestaProtocol/vesta/x/funding/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	keeper.HandleTokens(ctx)
}
