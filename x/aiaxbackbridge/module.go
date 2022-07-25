package aiax

import (
	"context"
	"encoding/json"

	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/keeper"
	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.AppModule      = AppModule{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec performs a no-op as the intrarelayer doesn't support Amino encoding
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (AppModuleBasic) ConsensusVersion() uint64 {
	return 1
}

func (AppModuleBasic) RegisterInterfaces(interfaceRegistry codectypes.InterfaceRegistry) {}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&types.GenesisState{})
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

func (AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {}

func (b AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(ctx))
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func NewAppModule(
	k keeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var gs types.GenesisState
	cdc.MustUnmarshalJSON(data, &gs)
	am.keeper.InitGenesis(ctx, gs)
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := types.GenesisState{}
	return cdc.MustMarshalJSON(&gs)
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (am AppModule) QuerierRoute() string {
	return ""
}

func (am AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (am AppModule) BeginBlock(ctx sdk.Context, rbb abci.RequestBeginBlock) {
	// Tried to deploy contract somewhere in InitGenesis but EVM does not deploys it outside of block
	// So it will be deployed in first block where it will not find stored address of contract
	// In future this allows easy updating (just deploy new contract and overwrite contract address)
	err := am.keeper.UpdateBackBridge(ctx)
	if err != nil {
		panic(err)
	}
}

func (am AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
