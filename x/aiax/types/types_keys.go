package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ModuleName = "aiax"
	TokenMain  = "aaiax"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

const (
	_ = byte(iota)

	/// Maps External ERC20 Token contract address to the Local ERC20 Token contract address.
	ExtERC20MappingKey

	/// Maps Local ERC20 Token contract address to the External Tokean contract address.
	LocalERC20MappingKey
)

var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

/// Make External ERC20 Token contract address to the Local ERC20 Token contract address key.
func MakeExtERC20MappingKey(extTokenAddress common.Address) []byte {
	return append([]byte{ExtERC20MappingKey}, extTokenAddress.Bytes()...)
}

/// Make Local ERC20 Token contract address to the External Tokean contract address key.
func MakeLocalERC20MappingKey(localTokenAddress common.Address) []byte {
	return append([]byte{LocalERC20MappingKey}, localTokenAddress.Bytes()...)
}
