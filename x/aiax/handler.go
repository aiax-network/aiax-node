package aiax

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/aiax-network/aiax-node/x/aiax/keeper"
	"github.com/aiax-network/aiax-node/x/aiax/types"
	// "github.com/peggyjv/gravity-bridge/module/x/gravity/keeper"
	// "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

func sendToEthereum(k keeper.Keeper, c context.Context, msg *types.MsgSendToEthereum) (*sdk.Result, error) {
	// ctx := sdk.UnwrapSDKContext(c)
	// sender, err := sdk.AccAddressFromBech32(msg.Sender)
	// if err != nil {
	// 	return nil, err
	// }

	// txID, err := k.GrvKeeper.createSendToEthereum(ctx, sender, msg.EthereumRecipient, msg.Amount, msg.BridgeFee)
	// if err != nil {
	// 	return nil, err
	// }

	// ctx.EventManager().EmitEvents([]sdk.Event{
	// 	sdk.NewEvent(
	// 		types.EventTypeBridgeWithdrawalReceived,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	// 		sdk.NewAttribute(types.AttributeKeyContract, k.GrvKeeper.getBridgeContractAddress(ctx)),
	// 		sdk.NewAttribute(types.AttributeKeyBridgeChainID, strconv.Itoa(int(k.GrvKeeper.GetParams(ctx).BridgeChainId))),
	// 		sdk.NewAttribute(types.AttributeKeyOutgoingTXID, strconv.Itoa(int(txID))),
	// 		sdk.NewAttribute(types.AttributeKeyNonce, fmt.Sprint(txID)),
	// 	),
	// 	sdk.NewEvent(
	// 		sdk.EventTypeMessage,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, msg.Type()),
	// 		sdk.NewAttribute(types.AttributeKeyOutgoingTXID, fmt.Sprint(txID)),
	// 	),
	// })

	// return &types.MsgSendToEthereumResponse{Id: txID}, nil

	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "___ not implemented %s message type: %T", types.ModuleName, msg)
}

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSendToEthereum:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "___ not implemented %s message type: %T", types.ModuleName, msg)
			// res, err := sendToEthereum(k, sdk.WrapSDKContext(ctx), msg)
			// return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
