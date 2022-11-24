package keeper

import (
	"vesta/x/vm/types"
)

var _ types.QueryServer = Keeper{}
