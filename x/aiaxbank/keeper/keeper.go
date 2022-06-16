package keeper

import (
	"fmt"

	"github.com/aiax-network/aiax-node/x/aiaxbank/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	acckeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	grvkeeper "github.com/peggyjv/gravity-bridge/module/x/gravity/keeper"
	"github.com/tendermint/tendermint/libs/log"
	evmkeeper "github.com/tharsis/ethermint/x/evm/keeper"
	irlkeeper "github.com/tharsis/evmos/x/intrarelayer/keeper"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramStore paramtypes.Subspace

	accKeeper *acckeeper.AccountKeeper
	banKeeper bankeeper.Keeper
	evmKeeper *evmkeeper.Keeper
	irlKeeper *irlkeeper.Keeper
	grvKeeper *grvkeeper.Keeper
}

func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	accKeeper *acckeeper.AccountKeeper,
	banKeeper bankeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	irlKeeper *irlkeeper.Keeper,
) Keeper {
	keeper := Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramStore: ps,
		accKeeper:  accKeeper,
		banKeeper:  banKeeper,
		evmKeeper:  evmKeeper,
		irlKeeper:  irlKeeper,
	}

	return keeper
}

func (k Keeper) AttachGravity(grvKeeper *grvkeeper.Keeper) {
	k.grvKeeper = grvKeeper
	k.grvKeeper.SetEthereumEventsHook(k)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
