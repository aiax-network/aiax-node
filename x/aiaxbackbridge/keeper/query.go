package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aiax-network/aiax-node/x/aiaxbackbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ types.QueryServer = Keeper{}
)

func (k Keeper) ContractInfo(ctx context.Context, req *types.ContractInfoRequest) (*types.ContractInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	uctx := sdk.UnwrapSDKContext(ctx)

	address := k.GetBackBridgeAddress(uctx)

	return &types.ContractInfoResponse{Address: address.Hex()}, nil
}
