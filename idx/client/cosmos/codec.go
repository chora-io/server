package cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
)

var ModuleBasics = []module.AppModuleBasic{
	groupmodule.AppModuleBasic{},
	bankmodule.AppModuleBasic{},
}

type Codec struct {
	*codec.ProtoCodec
}

func CustomCodec() *Codec {
	c := Codec{}

	// create interface registry
	ir := codectypes.NewInterfaceRegistry()

	// register types, crypto, and tx interfaces
	std.RegisterInterfaces(ir)

	// register module basic interfaces
	modBasic := module.NewBasicManager(ModuleBasics...)
	modBasic.RegisterInterfaces(ir)

	// create proto codec
	pc := codec.NewProtoCodec(ir)

	// set proto codec
	c.ProtoCodec = pc

	return &c
}

func (c Codec) Name() string {
	return "custom"
}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	switch x := v.(type) {
	case codec.ProtoMarshaler:
		return c.ProtoCodec.Marshal(x)
	default:
		return nil, fmt.Errorf("cannot marshal type %T", v)
	}
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	switch x := v.(type) {
	case codec.ProtoMarshaler:
		return c.ProtoCodec.Unmarshal(data, x)
	default:
		return fmt.Errorf("cannot unmarshal type %T", v)
	}
}
