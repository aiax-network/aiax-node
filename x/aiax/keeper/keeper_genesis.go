package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	log := ctx.Logger()
	k.SetParams(ctx, data.Params)

	// Get or create module account
	acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleName)
	if acc == nil {
		panic("Failed to init aiax-node module account")
	}
	log.Info("Module account Ethereum address: " + types.ModuleAddress.String())
	log.Info("Module account Cosmos address: " + sdk.AccAddress(types.ModuleAddress.Bytes()).String())
	address := data.Params.AiaxTokenContractAddress
	k.grvKeeper.SetCosmosOriginatedDenomToERC20(ctx, types.ModuleName, address)
  k.grvKeeper.SetCosmosOriginatedMintableStatus(ctx, address, true)
}
