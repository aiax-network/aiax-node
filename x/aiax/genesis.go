package aiax

import (
	"github.com/aiax-network/aiax-node/x/aiax/keeper"
	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	k.InitGenesis(ctx, data)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
