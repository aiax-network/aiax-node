package keeper

import (
	"fmt"

	"github.com/aiax-network/aiax-node/x/aiaxbank/types"
	"github.com/aiax-network/aiax-node/x/aiaxbank/types/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

// TODO: maybe move to separate module?

func (k Keeper) HandleEthereumEvent(ctx sdk.Context, eve gtypes.EthereumEvent) (bool, error) {
	switch event := eve.(type) {
	case *gtypes.SendToCosmosEvent:
		return k.handleSendToCosmosEvent(ctx, *event)
	}

	return false, nil
}

func (k Keeper) createCoinMetadata(ctx sdk.Context, evt gtypes.SendToCosmosEvent) (*banktypes.Metadata, error) {
	denom := types.CreateDenom("eth", evt.TokenContract)
	_, found := k.banKeeper.GetDenomMetaData(ctx, denom)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrTokenMapping, "token denom meta is already registered for: %s", evt.TokenContract)
	}

	metadata := banktypes.Metadata{
		Description: types.CreateDenomDescription(evt.Symbol, evt.Name, evt.TokenContract),
		Base:        denom,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    denom,
				Exponent: 0,
			},
			{
				Denom:    evt.Symbol,
				Exponent: evt.Decimals,
			},
		},
		Name:    denom,
		Symbol:  evt.Symbol,
		Display: evt.Symbol,
	}

	if err := metadata.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(err, "ERC20 token data is invalid for contract: %+v", evt)
	}
	k.banKeeper.SetDenomMetaData(ctx, metadata)

	return &metadata, nil
}

func (k Keeper) deployLocalERC20Contract(ctx sdk.Context, meta *banktypes.Metadata) (common.Address, error) {
	log := k.Logger(ctx)
	// Create constructor arguments with Name Symbol and Decimals (line 32 in this file) from event in gravity bridge
	decimals := uint8(meta.DenomUnits[1].Exponent)
	ctorArgs, err := contracts.ERC20BurnableAndMintableContract.ABI.Pack("", meta.Name, meta.Symbol, decimals)
	if err != nil {
		err = sdkerrors.Wrapf(err, "coin metadata is invalid %s", meta.Name)
		log.Error(err.Error())
		return common.Address{}, err
	}
	data := make([]byte, len(contracts.ERC20BurnableAndMintableContract.Bin)+len(ctorArgs))
	copy(data[:len(contracts.ERC20BurnableAndMintableContract.Bin)], contracts.ERC20BurnableAndMintableContract.Bin)
	copy(data[len(contracts.ERC20BurnableAndMintableContract.Bin):], ctorArgs)

	nonce, err := k.accKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		log.Error(err.Error())
		return common.Address{}, err
	}
	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.irlKeeper.CallEVMWithPayload(ctx, types.ModuleAddress, nil, data)
	if err != nil {
		err = fmt.Errorf("failed to deploy contract for %s", meta.Name)
		log.Error(err.Error())
		return common.Address{}, err
	}
	return contractAddr, nil
}

func (k Keeper) handleSendToCosmosEvent(ctx sdk.Context, evt gtypes.SendToCosmosEvent) (bool, error) {
	log := k.Logger(ctx)
	log.Info(fmt.Sprintf("Send to cosmos event: %+v", evt))

	extAddress := common.HexToAddress(evt.TokenContract)

	// TODO: Scan for all native tokens
	aiaxContract, exists := k.grvKeeper.GetCosmosOriginatedERC20(ctx, types.TokenMain)
	if exists && extAddress == aiaxContract {
		log.Info(fmt.Sprintf("Native Aiax token transfer to %s", evt.CosmosReceiver))
		return false, nil
	}

	aiaxContract, exists = k.grvKeeper.GetCosmosOriginatedERC20(ctx, "eth/"+extAddress.Hex())
	if exists && extAddress == aiaxContract {
		log.Info(fmt.Sprintf("Matched ERC20 (%s) token transfer to %s", "eth/"+extAddress.Hex(), evt.CosmosReceiver))
		return false, nil
	}

	meta, err := k.createCoinMetadata(ctx, evt)
	if err != nil {
		return false, err
	}

	localAddress, err := k.deployLocalERC20Contract(ctx, meta)
	if err != nil {
		return false, err
	}

	k.BindExternalAndLocalTokens(ctx, extAddress, localAddress)
	k.grvKeeper.SetCosmosOriginatedDenomToERC20(ctx, "eth/"+extAddress.Hex(), extAddress.Hex())
	k.grvKeeper.SetCosmosOriginatedMintableStatus(ctx, extAddress.Hex(), true)

	log.Info(fmt.Sprintf("Created ERC20 token for eth/%s", extAddress.Hex()))

	return false, nil
}
