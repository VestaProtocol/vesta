package keeper

import (
	"github.com/VestaProtocol/vesta/x/funding/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.ValidTokens(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ValidTokens returns the ValidTokens param
func (k Keeper) ValidTokens(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyValidTokens, &res)
	return
}
