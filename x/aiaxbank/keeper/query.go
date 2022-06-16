package keeper

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aiax-network/aiax-node/x/aiaxbank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ types.QueryServer = Keeper{}
)

func (k Keeper) DenomRepresentation(ctx context.Context, req *types.QueryDenomRepresentationRequest) (*types.QueryDenomRepresentationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	uctx := sdk.UnwrapSDKContext(ctx)

	if strings.HasPrefix(req.Denom, "eth/") {
		external := common.HexToAddress(req.Denom[4:])
		exists, internal := k.ExternalERC20LocalLookup(uctx, external)
		if exists {
			return &types.QueryDenomRepresentationResponse{
				ExternalAddress: external.Hex(),
				InternalAddress: internal.Hex(),
			}, nil
		}
	}

	// TODO: uncomment, it panics with "runtime error: invalid memory address or nil pointer dereference"
	// external, exists := k.grvKeeper.GetCosmosOriginatedERC20(uctx, req.Denom)
	// if exists {
	// 	return &types.QueryDenomRepresentationResponse{ExternalAddress: external.Hex()}, nil
	// }

	return &types.QueryDenomRepresentationResponse{}, nil
}
