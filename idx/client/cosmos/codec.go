package cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"

	authmodule "github.com/cosmos/cosmos-sdk/x/auth"
	vestingmodule "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank"
	distrmodule "github.com/cosmos/cosmos-sdk/x/distribution"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	govmodule "github.com/cosmos/cosmos-sdk/x/gov"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking"
	//
	// TODO: add ibc modules / update cosmos sdk version
	//ica "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts"
	//ibcfee "github.com/cosmos/ibc-go/v6/modules/apps/29-fee"
	//ibctransfer "github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	//ibc "github.com/cosmos/ibc-go/v6/modules/core"
	//
	// TODO: add regen modules / update cosmos sdk version
	//datamodule "github.com/regen-network/regen-ledger/x/data/v2/module"
	//ecocreditmodule "github.com/regen-network/regen-ledger/x/ecocredit/v3/module"
	//intertxmodule "github.com/regen-network/regen-ledger/x/intertx/module"
	//
	// TODO: add chora modules / update cosmos sdk version
	//contentmodule "github.com/choraio/mods/content/module"
	//geonodemodule "github.com/choraio/mods/geonode/module"
	//vouchermodule "github.com/choraio/mods/voucher/module"
)

var ModuleBasics = []module.AppModuleBasic{
	// sdk modules
	authmodule.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	bankmodule.AppModuleBasic{},
	distrmodule.AppModuleBasic{},
	feegrantmodule.AppModuleBasic{},
	govmodule.AppModuleBasic{},
	groupmodule.AppModuleBasic{},
	stakingmodule.AppModuleBasic{},
	vestingmodule.AppModuleBasic{},

	// TODO: add ibc modules / update cosmos sdk version
	//ibc.AppModuleBasic{},
	//ibcfee.AppModuleBasic{},
	//ibctransfer.AppModuleBasic{},
	//ica.AppModuleBasic{},

	// TODO: add regen modules / update cosmos sdk version
	//datamodule.Module{},
	//intertxmodule.AppModule{},
	//ecocreditmodule.Module{},

	// TODO: add chora modules / update cosmos sdk version
	//contentmodule.Module{},
	//geonodemodule.Module{},
	//vouchermodule.Module{},
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
