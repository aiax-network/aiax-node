package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ModuleName = "aiax"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

var ModuleAdress common.Address

func init() {
	ModuleAdress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

// TODO: ERC20 tokens mapping keys
