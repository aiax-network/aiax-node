package contracts

import (
	_ "embed"
	"encoding/json"

	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

var (
	//go:embed ERC20MinterBurner.json
	ERC20BurnableAndMintableJSON []byte // nolint: golint

	// ERC20BurnableAndMintableContract is the compiled erc20 contract
	ERC20BurnableAndMintableContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(ERC20BurnableAndMintableJSON, &ERC20BurnableAndMintableContract)
	if err != nil {
		panic(err)
	}
	if len(ERC20BurnableAndMintableContract.Bin) == 0 {
		panic("Failed to load embedded ERC20MinterBurner.json contract")
	}
}
