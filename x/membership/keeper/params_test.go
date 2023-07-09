package keeper_test

import (
	"testing"

	testkeeper "github.com/noria-net/module-membership/testutil/keeper"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MembershipKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
