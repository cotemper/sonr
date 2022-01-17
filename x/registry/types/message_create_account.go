package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAccount = "create_account"

var _ sdk.Msg = &MsgCreateAccount{}

func NewMsgCreateAccount(creator string, home string, deviceId string, fingerprint string, os string, model string, arch string, publicKey string, metadata string) *MsgCreateAccount {
	return &MsgCreateAccount{
		Creator:     creator,
		Home:        home,
		DeviceId:    deviceId,
		Fingerprint: fingerprint,
		Os:          os,
		Model:       model,
		Arch:        arch,
		PublicKey:   publicKey,
		Metadata:    metadata,
	}
}

func (msg *MsgCreateAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateAccount) Type() string {
	return TypeMsgCreateAccount
}

func (msg *MsgCreateAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}