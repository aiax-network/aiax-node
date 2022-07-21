package contracts

import (
	_ "embed"
	"encoding/json"

	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

var (
	//go:embed BackBridge.json
	BackBridgeJSON []byte // nolint: golint

	// BackBridgeContract is the compiled back bridge contract
	BackBridgeContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(BackBridgeJSON, &BackBridgeContract)
	if err != nil {
		panic(err)
	}
	if len(BackBridgeContract.Bin) == 0 {
		panic("Failed to load embedded ERC20MinterBurner.json contract")
	}
}
