package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) bindExternalAndLocalTokens(ctx sdk.Context, externaAddress common.Address, localAddress common.Address) {
	store := ctx.KVStore(k.storeKey)
	extKey := types.MakeExtERC20MappingKey(externaAddress)
	localKey := types.MakeLocalERC20MappingKey(localAddress)
	store.Set(extKey, localAddress.Bytes())
	store.Set(localKey, externaAddress.Bytes())
}

func (k Keeper) ExternalERC20LocalLookup(ctx sdk.Context, externalAddress common.Address) (bool, common.Address) {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.MakeExtERC20MappingKey(externalAddress))
	if bytes == nil {
		return false, common.Address{}
	}
	return true, common.BytesToAddress(bytes)
}

func (k Keeper) LocalERC20ExternalLookup(ctx sdk.Context, localAddress common.Address) (bool, common.Address) {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.MakeLocalERC20MappingKey(localAddress))
	if bytes == nil {
		return false, common.Address{}
	}
	return true, common.BytesToAddress(bytes)
}
