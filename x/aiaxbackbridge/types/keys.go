package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ModuleName = "aiaxbackbridge"
	StoreKey   = "hsakdjfsjksrgb"
	RouterKey  = ModuleName
)

const (
	_ = byte(iota)

	/// Stores actual address of back bridge
	BackBridgeAddr
)

var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}
