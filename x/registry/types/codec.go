package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterService{}, "registry/RegisterService", nil)
	cdc.RegisterConcrete(&MsgRegisterName{}, "registry/RegisterName", nil)
	cdc.RegisterConcrete(&MsgAccessName{}, "registry/AccessName", nil)
	cdc.RegisterConcrete(&MsgUpdateName{}, "registry/UpdateName", nil)
	cdc.RegisterConcrete(&MsgAccessService{}, "registry/AccessService", nil)
	cdc.RegisterConcrete(&MsgUpdateService{}, "registry/UpdateService", nil)
	cdc.RegisterConcrete(&MsgCreateWhois{}, "registry/CreateWhois", nil)
	cdc.RegisterConcrete(&MsgUpdateWhois{}, "registry/UpdateWhois", nil)
	cdc.RegisterConcrete(&MsgDeleteWhois{}, "registry/DeleteWhois", nil)
	cdc.RegisterConcrete(&MsgCreateWhatis{}, "registry/CreateWhatis", nil)
	cdc.RegisterConcrete(&MsgUpdateWhatis{}, "registry/UpdateWhatis", nil)
	cdc.RegisterConcrete(&MsgDeleteWhatis{}, "registry/DeleteWhatis", nil)
	cdc.RegisterConcrete(&MsgCreateThereis{}, "registry/CreateThereis", nil)
	cdc.RegisterConcrete(&MsgUpdateThereis{}, "registry/UpdateThereis", nil)
	cdc.RegisterConcrete(&MsgDeleteThereis{}, "registry/DeleteThereis", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhois{},
		&MsgUpdateWhois{},
		&MsgDeleteWhois{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhatis{},
		&MsgUpdateWhatis{},
		&MsgDeleteWhatis{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateThereis{},
		&MsgUpdateThereis{},
		&MsgDeleteThereis{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
