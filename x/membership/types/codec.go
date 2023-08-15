package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgEnroll{}, "membership/Enroll", nil)
	cdc.RegisterConcrete(&MsgUpdateStatus{}, "membership/UpdateStatus", nil)
	cdc.RegisterConcrete(&MsgUpdateDirectDemocracy{}, "membership/UpdateDirectDemocracy", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnroll{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateStatus{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateDirectDemocracy{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
