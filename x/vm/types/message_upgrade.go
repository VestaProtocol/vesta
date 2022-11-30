package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpgrade = "upgrade"

var _ sdk.Msg = &MsgUpgrade{}

func NewMsgUpgrade(creator string, contract string, code string) *MsgUpgrade {
	return &MsgUpgrade{
		Creator:  creator,
		Contract: contract,
		Code:     code,
	}
}

func (msg *MsgUpgrade) Route() string {
	return RouterKey
}

func (msg *MsgUpgrade) Type() string {
	return TypeMsgUpgrade
}

func (msg *MsgUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
