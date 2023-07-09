package keeper

import (
	"github.com/noria-net/module-membership/x/membership/types"
)

var _ types.QueryServer = Keeper{}
