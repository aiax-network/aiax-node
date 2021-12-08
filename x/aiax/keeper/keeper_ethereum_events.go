package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

func (k Keeper) HandleEthereumEvent(ctx sdk.Context, eve gtypes.EthereumEvent) (bool, error) {
	switch event := eve.(type) {
	case *gtypes.SendToCosmosEvent:
		return k.handleSendToCosmosEvent(ctx, *event)
	}

	return false, nil
}
