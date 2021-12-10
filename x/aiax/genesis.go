package aiax

import (
	"github.com/aiax-network/aiax-node/x/aiax/keeper"
	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accKeeper authkeeper.AccountKeeper,
	data types.GenesisState,
) {
	log := ctx.Logger()
	k.SetParams(ctx, data.Params)
	// Get or create module account
	acc := accKeeper.GetModuleAccount(ctx, types.ModuleName)
	if acc == nil {
		panic("Failed to init aiax-node module account")
	}
	log.Info("Module account Ethereum address: " + types.ModuleAddress.String())
	log.Info("Module account Cosmos address: " + sdk.AccAddress(types.ModuleAddress.Bytes()).String())
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
