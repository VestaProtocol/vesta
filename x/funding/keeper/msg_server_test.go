package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/x/funding/keeper"
	"github.com/VestaProtocol/vesta/x/funding/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// nolint
func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FundingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
