package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/TheMarstonConnell/vesta/testutil/keeper"
	"github.com/TheMarstonConnell/vesta/x/vm/keeper"
	"github.com/TheMarstonConnell/vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//nolint:all
func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
