package keeper

import (
	"math/big"

	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

// TODO: Return error instead of continuing (we know that this contract should be used only in this way)

func (k Keeper) PostTxProcessing(ctx sdk.Context, txHash common.Hash, logs []*ethtypes.Log) error {
	abi := contracts.BackBridgeContract.ABI
	addr := k.GetBackBridgeAddress(ctx)

	for _, i := range logs {
		if i.Address != addr {
			continue
		}

		eventID := i.Topics[0]

		event, err := abi.EventByID(eventID)
		if err != nil {
			return sdkerrors.Wrapf(err, "EVM contract event not found in ABI")
		}

		switch event.Name {
		case "MsgSendToEthereum":
			uevent, err := abi.Unpack(event.Name, i.Data)
			if err != nil {
				return sdkerrors.Wrapf(err, "Can't unpack event data")
			}

			from := common.BytesToAddress(i.Topics[1].Bytes())
			to := common.BytesToAddress(i.Topics[2].Bytes())

			amount_addr := common.BytesToAddress(i.Topics[3].Bytes())
			fee_addr := uevent[1].(common.Address)
			if amount_addr != fee_addr {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Amount and fee coins are not same")
			}

			denom, found := k.grvKeeper.GetCosmosOriginatedDenom(ctx, amount_addr.Hex())
			if !found {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Coins don't exists in gravity")
			}

			_, err = k.grvMsgServer.SendToEthereum(
				sdk.WrapSDKContext(ctx),
				gtypes.NewMsgSendToEthereum(
					sdk.AccAddress(from.Bytes()),
					to.Hex(),
					sdk.Coin{
						Denom:  denom,
						Amount: sdk.NewIntFromBigInt(uevent[0].(*big.Int)),
					},
					sdk.Coin{
						Denom:  denom,
						Amount: sdk.NewIntFromBigInt(uevent[2].(*big.Int)),
					},
				),
			)
			if err != nil {
				return sdkerrors.Wrapf(err, "Can't process gravity SendToEthereum")
			}

			// TODO: save response as event
			// May be useful to find SendToEthereum request by id and cancel it
		case "MsgRequestBatchTx":
			// TODO: maybe remove this from public contract interface since batches are sent periodically on their own
			if len(i.Data) > 0 {
				return sdkerrors.Wrapf(sdkerrors.ErrLogic, "Can't unpack event data")
			}

			denom_addr := common.BytesToAddress(i.Topics[1].Bytes())
			from := common.BytesToAddress(i.Topics[2].Bytes())

			denom, found := k.grvKeeper.GetCosmosOriginatedDenom(ctx, denom_addr.Hex())
			if !found {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Coins don't exists in gravity")
			}

			// Be carefully, this function contains nil pointer exception
			// So sometimes transaction fails
			// Good that this message is not necessary
			_, err = k.grvMsgServer.RequestBatchTx(
				sdk.WrapSDKContext(ctx),
				gtypes.NewMsgRequestBatchTx(
					denom,
					sdk.AccAddress(from.Bytes()),
				),
			)
			if err != nil {
				return sdkerrors.Wrapf(err, "Can't process gravity RequestBatchTx")
			}
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "Unknown event emitted")
		}
	}
	return nil
}
