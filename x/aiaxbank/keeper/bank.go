package keeper

import (
	"strings"

	"github.com/aiax-network/aiax-node/x/aiaxbank/types"
	"github.com/aiax-network/aiax-node/x/aiaxbank/types/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: support "eth/" and "aiax/"
// TODO: go over all 'amt' values (for all functions)
// TODO: safety checks as in banKeeper (for all functions)

func (k Keeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	return k.banKeeper.GetSupply(ctx, denom)
}

func (k Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	if strings.HasPrefix(amt[0].Denom, "eth/") {
		contract := common.HexToAddress(amt[0].Denom[4:])
		exists, contract := k.ExternalERC20LocalLookup(ctx, contract)

		if exists {
			senderAcc := k.accKeeper.GetModuleAccount(ctx, senderModule)
			if senderAcc == nil {
				panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", senderModule))
			}

			localAddr := senderAcc.GetAddress()
			sender := common.BytesToAddress(localAddr.Bytes())

			receiver := common.BytesToAddress(recipientAddr.Bytes())

			_, err := k.irlKeeper.CallEVM(
				ctx, contracts.ERC20BurnableAndMintableContract.ABI, sender,
				contract, "transfer", receiver, amt[0].Amount.BigInt())
			// TODO: does not fail on non existing contract
			return err
		}
	}

	return k.banKeeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	return k.banKeeper.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}

func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	if strings.HasPrefix(amt[0].Denom, "eth/") {
		contract := common.HexToAddress(amt[0].Denom[4:])
		exists, contract := k.ExternalERC20LocalLookup(ctx, contract)

		if exists {
			senderAcc := k.accKeeper.GetModuleAccount(ctx, recipientModule)
			if senderAcc == nil {
				panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", recipientModule))
			}

			localAddr := senderAcc.GetAddress()
			sender := common.BytesToAddress(localAddr.Bytes())

			receiver := common.BytesToAddress(senderAddr.Bytes())

			_, err := k.irlKeeper.CallEVM(
				ctx, contracts.ERC20BurnableAndMintableContract.ABI, receiver,
				contract, "transfer", sender, amt[0].Amount.BigInt())
			// TODO: does not fail on non existing contract
			return err
		}
	}
	return k.banKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

func (k Keeper) MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error {
	if strings.HasPrefix(amt[0].Denom, "eth/") {
		contract := common.HexToAddress(amt[0].Denom[4:])
		exists, contract := k.ExternalERC20LocalLookup(ctx, contract)

		if exists {
			recipientAcc := k.accKeeper.GetModuleAccount(ctx, name)
			if recipientAcc == nil {
				panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", name))
			}

			localAddr := recipientAcc.GetAddress()
			receiver := common.BytesToAddress(localAddr.Bytes())

			_, err := k.irlKeeper.CallEVM(
				ctx, contracts.ERC20BurnableAndMintableContract.ABI, types.ModuleAddress,
				contract, "mint", receiver, amt[0].Amount.BigInt())
			return err
		}
	}

	return k.banKeeper.MintCoins(ctx, name, amt)
}

func (k Keeper) BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error {
	if strings.HasPrefix(amt[0].Denom, "eth/") {
		contract := common.HexToAddress(amt[0].Denom[4:])
		exists, contract := k.ExternalERC20LocalLookup(ctx, contract)

		if exists {
			recipientAcc := k.accKeeper.GetModuleAccount(ctx, name)
			if recipientAcc == nil {
				panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", name))
			}

			localAddr := recipientAcc.GetAddress()
			receiver := common.BytesToAddress(localAddr.Bytes())

			_, err := k.irlKeeper.CallEVM(
				ctx, contracts.ERC20BurnableAndMintableContract.ABI, types.ModuleAddress,
				contract, "burn", receiver, amt[0].Amount.BigInt())
			return err
		}
	}

	return k.banKeeper.BurnCoins(ctx, name, amt)
}

func (k Keeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.banKeeper.GetAllBalances(ctx, addr)
}

func (k Keeper) GetDenomMetaData(ctx sdk.Context, denom string) (bank.Metadata, bool) {
	return k.banKeeper.GetDenomMetaData(ctx, denom)
}
