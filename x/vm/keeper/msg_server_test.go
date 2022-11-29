package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/x/vm/keeper"
	"github.com/VestaProtocol/vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//nolint:all
func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
