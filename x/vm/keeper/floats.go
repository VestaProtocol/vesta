package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dop251/goja"
)

const (
	numfield = "_num"
)

func (k Keeper) createDec(num string, ctx sdk.Context, vm *goja.Runtime) goja.Value {
	dec := vm.NewObject()
	nm, err := sdk.NewDecFromStr(num)
	if err != nil {
		ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", num))
		return goja.Undefined()
	}

	err = dec.Set(numfield, nm)
	if err != nil {
		ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", nm))
		return goja.Undefined()
	}

	err = dec.Set("add", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		right := call.Argument(0).ToObject(vm).Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}
		r, err := sdk.NewDecFromStr(right)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", right))
			return goja.Undefined()
		}

		s := l.Add(r).String()

		return k.createDec(s, ctx, vm)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("sub", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		right := call.Argument(0).ToObject(vm).Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}
		r, err := sdk.NewDecFromStr(right)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", right))
			return goja.Undefined()
		}

		s := l.Sub(r).String()

		return k.createDec(s, ctx, vm)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("mul", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		right := call.Argument(0).ToObject(vm).Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}
		r, err := sdk.NewDecFromStr(right)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", right))
			return goja.Undefined()
		}

		s := l.Mul(r).String()

		return k.createDec(s, ctx, vm)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("div", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		right := call.Argument(0).ToObject(vm).Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}
		r, err := sdk.NewDecFromStr(right)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", right))
			return goja.Undefined()
		}

		s := l.Quo(r).String()

		return k.createDec(s, ctx, vm)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("floor", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}

		s := l.TruncateDec().String()

		return k.createDec(s, ctx, vm)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("equal", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		right := call.Argument(0).ToObject(vm).Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}
		r, err := sdk.NewDecFromStr(right)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", right))
			return goja.Undefined()
		}

		s := l.Equal(r)

		val := vm.ToValue(s)

		return val
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("toString", func(call goja.FunctionCall) goja.Value {
		return dec.Get(numfield)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	err = dec.Set("toInt", func(call goja.FunctionCall) goja.Value {
		left := dec.Get(numfield).String()
		l, err := sdk.NewDecFromStr(left)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("Cannot parse %s as decimal", left))
			return goja.Undefined()
		}

		return vm.ToValue(l.TruncateInt().Int64())
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return goja.Undefined()
	}

	return dec
}

func (k Keeper) initFloats(ctx sdk.Context, vm *goja.Runtime, std *goja.Object) {
	err := std.Set("NewDec", func(call goja.FunctionCall) goja.Value {
		num := call.Argument(0).String()

		return k.createDec(num, ctx, vm)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
}
