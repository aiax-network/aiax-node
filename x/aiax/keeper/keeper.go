package keeper

import (
	"fmt"

	"github.com/aiax-network/aiax-node/x/aiax/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	acckeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	evmkeeper "github.com/tharsis/ethermint/x/evm/keeper"
  grvkeeper "github.com/peggyjv/gravity-bridge/module/x/gravity/keeper" 
	irlkeeper "github.com/tharsis/evmos/x/intrarelayer/keeper"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accKeeper acckeeper.AccountKeeperI
	banKeeper *bankeeper.Keeper
	evmKeeper *evmkeeper.Keeper
  grvKeeper *grvkeeper.Keeper
	irlKeeper *irlkeeper.Keeper
}

func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	accKeeper acckeeper.AccountKeeperI,
	banKeeper *bankeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
  grvKeeper *grvkeeper.Keeper,
	irlKeeper *irlkeeper.Keeper,
) Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(paramtypes.NewKeyTable())
	}
  keeper := Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramstore: ps,
		accKeeper:  accKeeper,
		banKeeper:  banKeeper,
		evmKeeper:  evmKeeper,
    grvKeeper:  grvKeeper,
		irlKeeper:  irlKeeper,
	}
  grvKeeper.SetEthereumEventsHook(keeper)

  return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
