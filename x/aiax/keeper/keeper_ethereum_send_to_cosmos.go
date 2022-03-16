package keeper

import (
	"fmt"

	"github.com/aiax-network/aiax-node/x/aiax/types"
	"github.com/aiax-network/aiax-node/x/aiax/types/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

func (k Keeper) createCoinMetadata(ctx sdk.Context, evt gtypes.SendToCosmosEvent) (*banktypes.Metadata, error) {
	_, found := k.banKeeper.GetDenomMetaData(ctx, types.CreateDenom(evt.TokenContract))
	if found {
		return nil, sdkerrors.Wrapf(types.ErrTokenMapping, "token denom meta is already registered for: %s", evt.TokenContract)
	}

	metadata := banktypes.Metadata{
		Description: types.CreateDenomDescription(evt.Symbol, evt.Name, evt.TokenContract),
		Base:        types.CreateDenom(evt.TokenContract),
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    types.CreateDenom(evt.TokenContract),
				Exponent: 0,
			},
			{
				Denom:    evt.Symbol,
				Exponent: evt.Decimals,
			},
		},
		Name:    types.CreateDenom(evt.TokenContract),
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
	ctorArgs, err := contracts.ERC20BurnableAndMintableContract.ABI.Pack("", meta.Name, meta.Symbol)
	if err != nil {
		err = sdkerrors.Wrapf(err, "coin metadata is invalid  %s", meta.Name)
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

	if evt.Native {
		log.Info(fmt.Sprintf("Native Aiax token transfer to %s", evt.CosmosReceiver))
    // TODO: Review
    k.grvKeeper.SetCosmosOriginatedDenomToERC20(ctx, types.TokenMain, evt.TokenContract)
		return false, nil
	}

	extAddress := common.HexToAddress(evt.TokenContract)
	exist, localAddress := k.ExternalERC20LocalLookup(ctx, extAddress)

	if !exist {
		meta, err := k.createCoinMetadata(ctx, evt)
		if err != nil {
			return false, err
		}
		localAddress, err = k.deployLocalERC20Contract(ctx, meta)
		if err != nil {
			return false, err
		}
		k.bindExternalAndLocalTokens(ctx, extAddress, localAddress)
		log.Info(fmt.Sprintf("Created ERC20 Contract Mapping %s => %s", extAddress.String(), localAddress.String()))
	} else {
		log.Info(fmt.Sprintf("ERC20 Contract Mapping %s => %s exists", extAddress.String(), localAddress.String()))
	}

	accAddress, err := sdk.AccAddressFromBech32(evt.CosmosReceiver)
	if err != nil {
		return false, err
	}
	receiver := common.BytesToAddress(accAddress.Bytes())

	// Mint required amount of local tokens
	_, err = k.irlKeeper.CallEVM(
		ctx, contracts.ERC20BurnableAndMintableContract.ABI, types.ModuleAddress,
		localAddress, "mint", receiver, evt.Amount.BigInt())
	if err != nil {
		return false, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeMintShadowERC20,
				sdk.NewAttribute(sdk.AttributeKeySender, evt.EthereumSender),
				sdk.NewAttribute(types.AttributeKeyCosmosReceiver, evt.CosmosReceiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, evt.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyERC20Address, evt.TokenContract),
				sdk.NewAttribute(types.AttributeKeyERC20LocalAddress, localAddress.String()),
				sdk.NewAttribute(types.AttributeKeyERC20Symbol, evt.Symbol),
			),
		},
	)

	log.Info(fmt.Sprintf("Minted %v on %s to %s", evt.Amount, accAddress.String(), localAddress))
	return true, nil
}
