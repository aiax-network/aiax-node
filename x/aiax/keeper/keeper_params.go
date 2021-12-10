package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/aiax-network/aiax-node/x/aiax/types"
)

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
  k.paramStore.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
  k.paramStore.GetParamSet(ctx, &params)
  return params
}
