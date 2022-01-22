package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWhois = "create_whois"
	TypeMsgUpdateWhois = "update_whois"
	TypeMsgDeleteWhois = "delete_whois"
)

var _ sdk.Msg = &MsgCreateWhois{}

func NewMsgCreateWhois(
	creator string,
	index string,
	did string,
	value string,

) *MsgCreateWhois {
	return &MsgCreateWhois{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgCreateWhois) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhois) Type() string {
	return TypeMsgCreateWhois
}

func (msg *MsgCreateWhois) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhois) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhois) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhois{}

func NewMsgUpdateWhois(
	creator string,
	index string,
	did string,
	value string,

) *MsgUpdateWhois {
	return &MsgUpdateWhois{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgUpdateWhois) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhois) Type() string {
	return TypeMsgUpdateWhois
}

func (msg *MsgUpdateWhois) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhois) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhois) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhois{}

func NewMsgDeleteWhois(
	creator string,
	index string,

) *MsgDeleteWhois {
	return &MsgDeleteWhois{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteWhois) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhois) Type() string {
	return TypeMsgDeleteWhois
}

func (msg *MsgDeleteWhois) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhois) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhois) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
