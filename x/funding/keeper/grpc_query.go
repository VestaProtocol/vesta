package keeper

import (
	"github.com/VestaProtocol/vesta/x/funding/types"
)

var _ types.QueryServer = Keeper{}
