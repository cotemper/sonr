package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWhatis = "create_whatis"
	TypeMsgUpdateWhatis = "update_whatis"
	TypeMsgDeleteWhatis = "delete_whatis"
)

var _ sdk.Msg = &MsgCreateWhatis{}

func NewMsgCreateWhatis(
	creator string,
	index string,
	did string,
	value string,

) *MsgCreateWhatis {
	return &MsgCreateWhatis{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgCreateWhatis) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhatis) Type() string {
	return TypeMsgCreateWhatis
}

func (msg *MsgCreateWhatis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhatis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhatis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhatis{}

func NewMsgUpdateWhatis(
	creator string,
	index string,
	did string,
	value string,

) *MsgUpdateWhatis {
	return &MsgUpdateWhatis{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgUpdateWhatis) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhatis) Type() string {
	return TypeMsgUpdateWhatis
}

func (msg *MsgUpdateWhatis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhatis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhatis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhatis{}

func NewMsgDeleteWhatis(
	creator string,
	index string,

) *MsgDeleteWhatis {
	return &MsgDeleteWhatis{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteWhatis) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhatis) Type() string {
	return TypeMsgDeleteWhatis
}

func (msg *MsgDeleteWhatis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhatis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhatis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
