package keeper

import (
	"github.com/VestaProtocol/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dop251/goja"
)

func (k Keeper) Cron(ctx sdk.Context) {
	jobs := k.GetAllCronjobs(ctx)

	for _, job := range jobs {

		if ctx.BlockHeight()%job.Interval > 0 {
			continue
		}

		program, found := k.GetProgram(ctx, job.Contract)
		if !found {
			k.RemoveCronjobs(ctx, job.Contract)
			continue
		}

		params := k.GetParams(ctx)
		denom := params.CronDenom

		var amt int64 = params.GetCronAmount()
		c := sdk.NewCoin(denom, sdk.NewInt(amt))
		cs := sdk.NewCoins(c)

		addr, err := sdk.AccAddressFromBech32(program.Address)
		if err != nil {
			k.RemoveCronjobs(ctx, job.Contract)
			continue
		}

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, cs)
		if err != nil {
			k.RemoveCronjobs(ctx, job.Contract)
			continue
		}

		err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, cs)
		if err != nil {
			k.RemoveCronjobs(ctx, job.Contract)
			continue
		}

		source, ok := k.GetContracts(ctx, GetContractVersion(program, "-1"))
		if !ok {
			k.RemoveCronjobs(ctx, job.Contract)
			continue
		}

		_, err = k.executeContract(ctx, job.Contract, source.Source, job.Function, nil, []goja.Value{})
		if err != nil {
			continue
		}
	}
}
