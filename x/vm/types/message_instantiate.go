package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInstantiate = "instantiate"

var _ sdk.Msg = &MsgInstantiate{}

func NewMsgInstantiate(creator string, name string, code string, args string) *MsgInstantiate {
	return &MsgInstantiate{
		Creator: creator,
		Name:    name,
		Code:    code,
		Args:    args,
	}
}

func (msg *MsgInstantiate) Route() string {
	return RouterKey
}

func (msg *MsgInstantiate) Type() string {
	return TypeMsgInstantiate
}

func (msg *MsgInstantiate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInstantiate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInstantiate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
