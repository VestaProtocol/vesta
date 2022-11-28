package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dop251/goja"
)

type Injector interface {
	Name() string
	Inject(ctx sdk.Context, creator sdk.AccAddress, contractName string, contractAddress sdk.AccAddress, vm *goja.Object, readonly bool) error
}
