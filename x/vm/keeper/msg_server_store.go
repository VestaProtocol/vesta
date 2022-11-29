package keeper

import (
	"context"
	"fmt"

	"github.com/VestaProtocol/vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sourceCode := msg.Source

	contract := types.Contracts{
		Id:     0,
		Source: sourceCode,
	}

	id := k.AppendContracts(ctx, contract)

	return &types.MsgStoreResponse{Code: fmt.Sprintf("%d", id)}, nil
}
