package keeper

import (
	"math/big"

	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	gtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
)

// TODO: Return error instead of continuing (we know that this contract should be used only in this way)

func (k Keeper) PostTxProcessing(ctx sdk.Context, txHash common.Hash, logs []*ethtypes.Log) error {
	log := k.Logger(ctx)

	abi := contracts.BackBridgeContract.ABI
	addr := k.GetBackBridgeAddress(ctx)

	for _, i := range logs {
		if i.Address != addr {
			continue
		}

		eventID := i.Topics[0]

		event, err := abi.EventByID(eventID)
		if err != nil {
			return err
		}

		// For compatibility on testnet update
		// TODO: Remove
		if ctx.ChainID() == "aiax_12344123324-1" && ctx.BlockHeight() < 120050 {
			_, err := abi.Unpack(event.Name, i.Data)
			if err != nil {
				return err
			}
		}

		switch event.Name {
		case "MsgSendToEthereum":
			uevent, err := abi.Unpack(event.Name, i.Data)
			if err != nil {
				return err
			}

			from := common.BytesToAddress(i.Topics[1].Bytes())
			to := common.BytesToAddress(i.Topics[2].Bytes())

			amount_addr := common.BytesToAddress(i.Topics[3].Bytes())
			fee_addr := uevent[1].(common.Address)
			if amount_addr != fee_addr {
				continue
			}

			denom, found := k.grvKeeper.GetCosmosOriginatedDenom(ctx, amount_addr.Hex())
			if !found {
				continue
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
				return err
			}

			// TODO: Save response as event
		case "MsgRequestBatchTx":
			denom_addr := common.BytesToAddress(i.Topics[1].Bytes())
			from := common.BytesToAddress(i.Topics[2].Bytes())

			denom, found := k.grvKeeper.GetCosmosOriginatedDenom(ctx, denom_addr.Hex())
			if !found {
				continue
			}

			_, err = k.grvMsgServer.RequestBatchTx(
				sdk.WrapSDKContext(ctx),
				gtypes.NewMsgRequestBatchTx(
					denom,
					sdk.AccAddress(from.Bytes()),
				),
			)
			if err != nil {
				return err
			}
		default:
			log.Info("emitted event", "name", event.Name, "signature", event.Sig)
		}
	}
	return nil
}
