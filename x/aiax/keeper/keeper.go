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
	irlkeeper "github.com/tharsis/evmos/x/intrarelayer/keeper"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accKeeper acckeeper.AccountKeeperI
	banKeeper *bankeeper.Keeper
	evmKeeper *evmkeeper.Keeper
	irlKeeper *irlkeeper.Keeper
}

func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	accKeeper acckeeper.AccountKeeperI,
	banKeeper *bankeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	irlKeeper *irlkeeper.Keeper,
) Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(paramtypes.NewKeyTable())
	}
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramstore: ps,
		accKeeper:  accKeeper,
		banKeeper:  banKeeper,
		evmKeeper:  evmKeeper,
		irlKeeper:  irlKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
