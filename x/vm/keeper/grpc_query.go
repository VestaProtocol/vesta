package keeper

import (
	"github.com/TheMarstonConnell/vesta/x/vm/types"
)

var _ types.QueryServer = Keeper{}
