package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiaxbank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	log := ctx.Logger()

	// Get or create module account
	acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleName)
	if acc == nil {
		panic("Failed to init aiax-node module account")
	}
	log.Info("Module aiaxbank account Ethereum address: " + types.ModuleAddress.String())
	log.Info("Module aiaxbank account Cosmos address: " + sdk.AccAddress(types.ModuleAddress.Bytes()).String())
}
