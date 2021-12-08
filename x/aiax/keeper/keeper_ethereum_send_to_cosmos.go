package keeper

import (
	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

func (k Keeper) createCoinMetadata(ctx sdk.Context, evt gtypes.SendToCosmosEvent) (*banktypes.Metadata, error) {

	_, found := k.banKeeper.GetDenomMetaData(ctx, types.CreateDenom(evt.TokenContract))
	if found {
    return nil, sdkerrors.Wrapf(types.ErrTokenMapping, "token denom meta is already registered for: %s", evt.TokenContract)
	}
  
  metadata := banktypes.Metadata{
    Description: types.CreateDenomDescription(evt.TokenContract),
    Base: types.CreateDenom(evt.TokenContract),
    DenomUnits: []*banktypes.DenomUnit{
      {
        Denom: types.CreateDenom(evt.TokenContract),
        Exponent: 0,
      },
      {
        Denom: evt.Name,
        Exponent: evt.Decimals,
      },
    },
    Name: types.CreateDenom(evt.TokenContract),
    Symbol: evt.Symbol,
    Display: evt.Name,
  }

  if err := metadata.Validate(); err != nil {
    return nil, sdkerrors.Wrapf(err, "ERC20 token data is invalid for contract: %+v", evt)
  }
  k.banKeeper.SetDenomMetaData(ctx, metadata)

  return &metadata, nil
}

func (k Keeper) handleSendToCosmosEvent(ctx sdk.Context, evt gtypes.SendToCosmosEvent) (bool, error) {

	extAddress := common.HexToAddress(evt.TokenContract)
	exist, localAddress := k.ExternalERC20LocalLookup(ctx, extAddress)

	if !exist {

	}

	// TODO mint local tokens

	// Sample send to cosmos event
	// event_nonce:14
	// token_contract:"0xB581C9264f59BF0289fA76D61B2D0746dCE3C30D"
	// amount:"10"
	// ethereum_sender:"0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
	// cosmos_receiver:"aiax1y4rte57ggcs7jakcrpdfr2fz4em7empscx5dqz"
	// ethereum_height:282
	// name:"Aiax test token one"
	// symbol:"TONE"
	// decimals:18
	return true, nil
}
