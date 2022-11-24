package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExecute = "execute"

var _ sdk.Msg = &MsgExecute{}

func NewMsgExecute(creator string, contract string, function string, args string) *MsgExecute {
	return &MsgExecute{
		Creator:  creator,
		Contract: contract,
		Function: function,
		Args:     args,
	}
}

func (msg *MsgExecute) Route() string {
	return RouterKey
}

func (msg *MsgExecute) Type() string {
	return TypeMsgExecute
}

func (msg *MsgExecute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExecute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
