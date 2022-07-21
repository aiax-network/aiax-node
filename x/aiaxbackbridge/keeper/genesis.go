package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleName)
	if acc == nil {
		panic("Failed to init aiaxbackbridge module account")
	}
}
