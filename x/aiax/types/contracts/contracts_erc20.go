package contracts

import (
	_ "embed"
	"encoding/json"

	"github.com/aiax-network/aiax-node/x/aiax/types"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

var (
	//go:embed ERC20MinterBurner.json
	ERC20BurnableAndMintableJSON []byte // nolint: golint

	// ERC20BurnableAndMintableContract is the compiled erc20 contract
	ERC20BurnableAndMintableContract evmtypes.CompiledContract

	// ERC20BurnableAndMintableAddress is the irm module address
	ERC20BurnableAndMintableAddress common.Address
)

func init() {
	ERC20BurnableAndMintableAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20BurnableAndMintableJSON, &ERC20BurnableAndMintableContract)
	if err != nil {
		panic(err)
	}
	if len(ERC20BurnableAndMintableContract.Bin) == 0 {
		panic("Failed to load embedded ERC20MinterBurner.json contract")
	}
}
