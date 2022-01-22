package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateThereis = "create_thereis"
	TypeMsgUpdateThereis = "update_thereis"
	TypeMsgDeleteThereis = "delete_thereis"
)

var _ sdk.Msg = &MsgCreateThereis{}

func NewMsgCreateThereis(
	creator string,
	index string,
	did string,
	value string,

) *MsgCreateThereis {
	return &MsgCreateThereis{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgCreateThereis) Route() string {
	return RouterKey
}

func (msg *MsgCreateThereis) Type() string {
	return TypeMsgCreateThereis
}

func (msg *MsgCreateThereis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateThereis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateThereis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateThereis{}

func NewMsgUpdateThereis(
	creator string,
	index string,
	did string,
	value string,

) *MsgUpdateThereis {
	return &MsgUpdateThereis{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgUpdateThereis) Route() string {
	return RouterKey
}

func (msg *MsgUpdateThereis) Type() string {
	return TypeMsgUpdateThereis
}

func (msg *MsgUpdateThereis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateThereis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateThereis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteThereis{}

func NewMsgDeleteThereis(
	creator string,
	index string,

) *MsgDeleteThereis {
	return &MsgDeleteThereis{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteThereis) Route() string {
	return RouterKey
}

func (msg *MsgDeleteThereis) Type() string {
	return TypeMsgDeleteThereis
}

func (msg *MsgDeleteThereis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteThereis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteThereis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
