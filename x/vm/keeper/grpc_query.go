package keeper

import (
	"github.com/VestaProtocol/vesta/x/vm/types"
)

var _ types.QueryServer = Keeper{}
