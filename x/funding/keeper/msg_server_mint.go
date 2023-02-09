package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/VestaProtocol/vesta/x/funding/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	c, err := sdk.ParseCoinNormalized(msg.Token)
	if err != nil {
		return nil, err
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	var tokens []string
	err = json.Unmarshal([]byte(k.GetParams(ctx).ValidTokens), &tokens)
	if err != nil {
		return nil, err
	}

	for _, t := range tokens {
		if c.Denom == t {
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(c))
			if err != nil {
				return nil, err
			}
			toMint := sdk.NewCoins(sdk.NewCoin("uvst", c.Amount))
			err = k.bankKeeper.MintCoins(ctx, types.ModuleName, toMint)
			if err != nil {
				return nil, err
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, toMint)
			if err != nil {
				return nil, err
			}
			return &types.MsgMintResponse{}, nil
		}
	}

	return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, fmt.Sprintf("could not parse %s", msg.Token))
}
