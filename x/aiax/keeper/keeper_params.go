package keeper

import (
	"strings"

	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
  params.AiaxTokenContractAddress = strings.ToLower(params.AiaxTokenContractAddress)
  k.paramStore.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
  k.paramStore.GetParamSet(ctx, &params)
  return params
}
