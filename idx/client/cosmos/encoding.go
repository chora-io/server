package cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
)

var ModuleBasics = []module.AppModuleBasic{
	groupmodule.AppModuleBasic{},
}

type GRPCCodec struct {
	*codec.ProtoCodec
	codectypes.InterfaceRegistry
}

func CustomCodec() *GRPCCodec {
	cm := GRPCCodec{}
	modBasic := module.NewBasicManager(ModuleBasics...)
	ir := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(ir)
	modBasic.RegisterInterfaces(ir)
	cdc := codec.NewProtoCodec(ir)
	cm.ProtoCodec = cdc
	cm.InterfaceRegistry = ir
	return &cm
}

func (c GRPCCodec) Name() string {
	return "custom"
}

func (c GRPCCodec) Marshal(v interface{}) ([]byte, error) {
	switch x := v.(type) {
	case codec.ProtoMarshaler:
		return c.ProtoCodec.Marshal(x)
	default:
		return nil, fmt.Errorf("cannot marshal type %T", v)
	}
}

func (c GRPCCodec) Unmarshal(data []byte, v interface{}) error {
	switch x := v.(type) {
	case codec.ProtoMarshaler:
		return c.ProtoCodec.Unmarshal(data, x)
	default:
		return fmt.Errorf("cannot unmarshal type %T", v)
	}
}
