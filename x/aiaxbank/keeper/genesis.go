package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiaxbank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Get or create module account
	acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleName)
	if acc == nil {
		panic("Failed to init aiaxbank module account")
	}
}
