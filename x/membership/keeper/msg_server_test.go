package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/noria-net/module-membership/testutil/keeper"
	"github.com/noria-net/module-membership/x/membership/keeper"
	"github.com/noria-net/module-membership/x/membership/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MembershipKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
