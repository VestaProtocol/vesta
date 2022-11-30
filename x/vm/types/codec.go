package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStore{}, "vm/Store", nil)
	cdc.RegisterConcrete(&MsgExecute{}, "vm/Execute", nil)
	cdc.RegisterConcrete(&MsgInstantiate{}, "vm/Instantiate", nil)
	cdc.RegisterConcrete(&MsgUpgrade{}, "vm/Upgrade", nil)
	cdc.RegisterConcrete(&MsgCron{}, "vm/Cron", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStore{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecute{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInstantiate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpgrade{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCron{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
