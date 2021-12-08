package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

func (k Keeper) HandleEthereumEvent(ctx sdk.Context, eve gtypes.EthereumEvent) (bool, error) {
	switch event := eve.(type) {
	case *gtypes.SendToCosmosEvent:
		ctx.Logger().Error(fmt.Sprintf("!!! SYMB %v", event))
	}

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
	return false, nil
}
