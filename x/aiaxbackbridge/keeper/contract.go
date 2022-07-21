package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types"
	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k Keeper) deployBackBridge(ctx sdk.Context) (common.Address, error) {
	log := k.Logger(ctx)

	ctorArgs, err := contracts.BackBridgeContract.ABI.Pack("")
	if err != nil {
		return common.Address{}, err
	}

	data := make([]byte, len(contracts.BackBridgeContract.Bin)+len(ctorArgs))
	copy(data[:len(contracts.BackBridgeContract.Bin)], contracts.BackBridgeContract.Bin)
	copy(data[len(contracts.BackBridgeContract.Bin):], ctorArgs)

	nonce, err := k.accKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		log.Error(err.Error())
		return common.Address{}, err
	}
	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.irlKeeper.CallEVMWithPayload(ctx, types.ModuleAddress, nil, data)
	if err != nil {
		return common.Address{}, err
	}
	return contractAddr, nil
}

func (k Keeper) UpdateBackBridge(ctx sdk.Context) error {
	log := k.Logger(ctx)

	store := ctx.KVStore(k.storeKey)

	bytes := store.Get([]byte{types.BackBridgeAddr})
	if bytes != nil {
		return nil
	}

	caddr, err := k.deployBackBridge(ctx)
	if err != nil {
		return err
	}
	log.Info("Deployed BackBridge.sol at " + caddr.Hex())

	store.Set([]byte{types.BackBridgeAddr}, caddr.Bytes())

	return nil
}

func (k Keeper) GetBackBridgeAddress(ctx sdk.Context) common.Address {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get([]byte{types.BackBridgeAddr})
	if bytes == nil {
		return common.Address{}
	}
	return common.BytesToAddress(bytes)
}
