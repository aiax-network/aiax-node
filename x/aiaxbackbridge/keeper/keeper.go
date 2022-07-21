package keeper

import (
	"fmt"

	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	acckeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	grvkeeper "github.com/peggyjv/gravity-bridge/module/x/gravity/keeper"
	grvtypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
	"github.com/tendermint/tendermint/libs/log"
	evmkeeper "github.com/tharsis/ethermint/x/evm/keeper"
	irlkeeper "github.com/tharsis/evmos/x/intrarelayer/keeper"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramStore paramtypes.Subspace

	accKeeper    *acckeeper.AccountKeeper
	banKeeper    bankeeper.Keeper
	evmKeeper    *evmkeeper.Keeper
	irlKeeper    *irlkeeper.Keeper
	grvKeeper    *grvkeeper.Keeper
	grvMsgServer grvtypes.MsgServer
}

func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	accKeeper *acckeeper.AccountKeeper,
	banKeeper bankeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	irlKeeper *irlkeeper.Keeper,
	grvKeeper *grvkeeper.Keeper,
) Keeper {
	keeper := Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramStore: ps,
		accKeeper:  accKeeper,
		banKeeper:  banKeeper,
		evmKeeper:  evmKeeper,
		irlKeeper:  irlKeeper,
		grvKeeper:  grvKeeper,
	}

	keeper.grvMsgServer = grvkeeper.NewMsgServerImpl(*grvKeeper)

	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
