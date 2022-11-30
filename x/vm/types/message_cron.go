package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCron = "cron"

var _ sdk.Msg = &MsgCron{}

func NewMsgCron(creator string, contract string, function string, interval string) *MsgCron {
	return &MsgCron{
		Creator:  creator,
		Contract: contract,
		Function: function,
		Interval: interval,
	}
}

func (msg *MsgCron) Route() string {
	return RouterKey
}

func (msg *MsgCron) Type() string {
	return TypeMsgCron
}

func (msg *MsgCron) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCron) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCron) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
